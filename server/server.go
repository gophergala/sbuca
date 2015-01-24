package server

import (
  "fmt"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  //"github.com/gophergala/sbuca/x509util"
  "net/http"
  "encoding/pem"
  "github.com/gophergala/sbuca/pkix"
  "github.com/gophergala/sbuca/ca"
)


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
    csr, err := pkix.NewCertificateRequestFromPEM([]byte(csrString))
    if err != nil {
      panic(err)
    }

    newCA, err := ca.NewCA(".")
    if err != nil {
      panic(err)
    }

    /*
    notAfter  := notBefore.Add(time.Hour*365*24)
    keyUsage  := x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature
    extKeyUsage := []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
    template := &x509.Certificate{
      SerialNumber: big.NewInt(1),
      Subject: csr.Csr.Subject,
      NotBefore: notBefore,
      NotAfter: notAfter,
      KeyUsage: keyUsage,
      ExtKeyUsage: extKeyUsage,
      BasicConstraintsValid: true,
    }

    der, err := x509.CreateCertificate(rand.Reader, template, caCert.Crt, caKey.PublicKey, caKey.PrivateKey)
    if err != nil {
      panic(err)
    }
    */
    cert, err := newCA.IssueCertificate(csr)
    if err != nil {
      panic(err)
    }

    pemB := &pem.Block{
      Type: "CERTIFICATE",
      Bytes: cert.DerBytes,
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
