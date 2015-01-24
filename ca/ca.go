package ca

import (
  "strconv"
  "time" 
  //"errors"
  "math/big"
  "io/ioutil"
  "strings"
  "github.com/gophergala/sbuca/pkix"
  "crypto/x509"
  "crypto/rand"
)


type CA struct {
  RootDir string
  Certificate *pkix.Certificate
  Key *pkix.Key
}
func NewCA(rootDir string) (*CA, error) {

  certificate, err := pkix.NewCertificateFromPEMFile(rootDir + "/ca/ca.crt")
  if err != nil {
    return nil, err
  }
  key, err := pkix.NewKeyFromPrivateKeyPEMFile(rootDir + "/ca/ca.key")
  if err != nil {
    return nil, err
  }
  newCA := &CA{
    RootDir: rootDir,
    Certificate: certificate,
    Key: key,
  }

  return newCA, nil
}
func (ca *CA) GetSerialNumber() (big.Int, error) {
  snStr, err := ioutil.ReadFile(ca.RootDir + "/ca/ca.srl")
  if err != nil {
    panic(err)
  }
  snInt, err := strconv.Atoi(strings.Trim(string(snStr), "\n"))
  if err != nil {
    panic(err)
  }
  sn := big.NewInt(int64(snInt))

  return *sn, nil
}
func (ca *CA) IncreaseSerialNumber() error {
  ioutil.ReadFile(ca.RootDir + "/ca.srl")
  snStr, err := ioutil.ReadFile(ca.RootDir + "/ca.srl")
  if err != nil {
    panic(err)
  }
  snInt, err := strconv.Atoi(strings.Trim(string(snStr), "\n"))
  if err != nil {
    panic(err)
  }
  nextSnInt := snInt + 1
  nextSnStr := strconv.Itoa(nextSnInt)
  ioutil.WriteFile(ca.RootDir + "/ca.srl", []byte(nextSnStr), 0600)

  return nil
}
func (ca *CA) IssueCertificate(csr *pkix.CertificateRequest) (*pkix.Certificate, error) {

  notBefore := time.Now()
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

  derBytes, err := x509.CreateCertificate(rand.Reader, template, ca.Certificate.Crt, ca.Key.PublicKey, ca.Key.PrivateKey)
  if err != nil {
    return nil, err
  }

  crt, err := pkix.NewCertificateFromDER(derBytes)
  if err != nil {
    return nil, err
  }
  return crt, nil
}
