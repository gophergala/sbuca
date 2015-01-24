package server

import (
  "fmt"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  //"github.com/gophergala/sbuca/x509util"
  "net/http"
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

    cert, err := newCA.IssueCertificate(csr)
    if err != nil {
      panic(err)
    }

    certPem, err := cert.ToPEM()
    if err != nil {
      panic(err)
    }
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
