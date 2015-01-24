package pkix

import (

  "errors"
  "crypto/x509"
  "io/ioutil"
  "encoding/pem"
)

type Certificate struct {

  DerBytes []byte

  Crt *x509.Certificate

}

func NewCertificateFromDER(derBytes []byte) (*Certificate, error) {

  crt, err := x509.ParseCertificate(derBytes)
  if err != nil {
    return nil, err
  }

  cert := &Certificate{
    DerBytes: derBytes,
    Crt: crt,
  }

  return cert, nil
}
func NewCertificateFromPEM(pemBytes []byte) (*Certificate, error) {

  pemBlock, _ := pem.Decode(pemBytes)
  if pemBlock == nil {
    return nil, errors.New("PEM decode failed")
  }

  crt, err := x509.ParseCertificate(pemBlock.Bytes) 
  if err != nil {
    return nil, err
  }

  cert := &Certificate{
    DerBytes: pemBlock.Bytes,
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

func (certificate *Certificate) ToPEM() ([]byte, error) {

  pemBlock := &pem.Block{
    Type: "CERTIFICATE",
    Bytes: certificate.DerBytes,
  }

  pemBytes := pem.EncodeToMemory(pemBlock)

  return pemBytes, nil
}
