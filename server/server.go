package server

import (
  "math/big"
  "time"
  "fmt"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/gophergala/sbuca/x509util"
  "net/http"
  "encoding/pem"
  "crypto/x509"
  "crypto/rand"
  "crypto/rsa"
)

func GetCACert() *x509.Certificate {
  cert, err := x509util.PemFileToCertificate("ca/ca.crt")
  if err != nil {
    panic(err)
  }
  return cert
}

func GetCAPrivateKey() *rsa.PrivateKey {
  key, err := x509util.PemFileToRsaPrivateKey("ca/ca.key")
  if err != nil {
    panic(err)
  }
  return key
}

func Run() {
  fmt.Print("start...")

  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", func() string {
    return "hello\n"
  })
  m.Get("/certificates/:id", func(params martini.Params, r render.Render) {
  })
  m.Get("/certificates", func(params martini.Params, r render.Render) {
    r.JSON(200, map[string]interface{}{
      "certificate": map[string]interface{}{
        "id": "1",
        "crt": "2",
      },
    })
  })
  m.Post("/certificates", func(req *http.Request, params martini.Params, r render.Render) {

    csrString := req.PostFormValue("csr")
    pemBlock, _ := pem.Decode([]byte(csrString))
    if pemBlock == nil {
      fmt.Println("pem decode failed")
    }
    csr, err := x509.ParseCertificateRequest(pemBlock.Bytes)
    if err != nil {
      panic(csr)
    }

    caCert := GetCACert()
    caPrivKey := GetCAPrivateKey()
    fmt.Println(caCert.Version)

    notBefore := time.Now()
    notAfter  := notBefore.Add(time.Hour*365*24)
    keyUsage  := x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature
    extKeyUsage := []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
    template := &x509.Certificate{
      SerialNumber: big.NewInt(1),
      Subject: csr.Subject,
      NotBefore: notBefore,
      NotAfter: notAfter,
      KeyUsage: keyUsage,
      ExtKeyUsage: extKeyUsage,
      BasicConstraintsValid: true,
    }

    der, err := x509.CreateCertificate(rand.Reader, template, caCert, caPrivKey.Public(), caPrivKey)
    if err != nil {
      panic(err)
    }

    //certOut, err := os.Create("cert.pem")
    //if err != nil {
    //  panic(err)
    //}

    pemB := &pem.Block{
      Type: "CERTIFICATE",
      Bytes: der,
    }
    certPem := pem.EncodeToMemory(pemB)
    /*
    r.JSON(200, map[string]interface{}{
      "certificate": map[string]interface{}{
        "id": "1",
        "crt": "2",
        "csr": csr,
      },
    })
    */
    r.Data(200, certPem)
  })

  m.Run()

}
