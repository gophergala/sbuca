package pkix

import (

  "errors"
  "crypto/x509"
  "io/ioutil"
  "encoding/pem"
)

type CertificateRequest struct {

  pemBytes []byte

  Csr *x509.CertificateRequest

}

func NewCertificateRequestFromPEM(pemBytes []byte) (*CertificateRequest, error) {

  pemBlock, _ := pem.Decode(pemBytes)
  if pemBlock == nil {
    return nil, errors.New("PEM decode failed")
  }

  csr, err := x509.ParseCertificateRequest(pemBlock.Bytes)
  if err != nil {
    return nil, err
  }

  certificateRequest := &CertificateRequest{
    pemBytes: pemBytes,
    Csr: csr,
  }

  return certificateRequest, nil
}
func NewCertificateRequestFromPEMFile(filename string) (*CertificateRequest, error) {

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return NewCertificateRequestFromPEM(data)
}




