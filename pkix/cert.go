package pkix

import (

  "errors"
  "crypto/x509"
  "io/ioutil"
  "encoding/pem"
)

type Certificate struct {

  derBytes []byte

  Crt *x509.Certificate

}

func NewCertificateFromPEM(derBytes []byte) (*Certificate, error) {

  pemBlock, _ := pem.Decode(derBytes)
  if pemBlock == nil {
    return nil, errors.New("PEM decode failed")
  }

  crt, err := x509.ParseCertificate(pemBlock.Bytes) 
  if err != nil {
    return nil, err
  }

  cert := &Certificate{
    derBytes: derBytes,
    Crt: crt,
  }

  return cert, nil
}
func NewCertificateFromPEMFile(filename string) (*Certificate, error) {

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return NewCertificateFromPEM(data)
}



