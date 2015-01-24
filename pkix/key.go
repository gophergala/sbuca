package pkix

import (
  "encoding/pem"
  "crypto"
  "io/ioutil"
  "errors"
  "crypto/x509"
  "crypto/rsa"
)

type Key struct {
  /*
  PublicKey *crypto.PublicKey
  PrivateKey *crypto.PrivateKey
  */
  PublicKey crypto.PublicKey
  PrivateKey *rsa.PrivateKey
}

func NewKeyFromPrivateKeyPEM(pemBytes []byte) (*Key, error) {
  // currently we only support rsa

  pemBlock, _ := pem.Decode(pemBytes)
  if pemBlock == nil {
    return nil, errors.New("decode pem failed")
  }

  privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
  if err != nil {
    return nil, err
  }

  newKey := &Key{
    PrivateKey: privateKey,
    PublicKey: privateKey.Public(),
  }

  return newKey, nil
}
func NewKeyFromPrivateKeyPEMFile(filename string) (*Key, error) {

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return NewKeyFromPrivateKeyPEM(data)

}
