package main

import (
  "strconv"
  //"errors"
  "fmt"
  "path/filepath"
  "os"
  "math/big"
  "io/ioutil"
  "strings"
)


type CA struct {
  RootDir string
}
func NewCA(rootDir string) *CA {
  return &CA{RootDir: rootDir}
}
func (ca *CA) GetSerialNumber() (big.Int, error) {
  snStr, err := ioutil.ReadFile(ca.RootDir + "/ca.srl")
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
//func (ca *CA) Init {
//}
//func (ca *CA) GetSerialNumber {
//}

func main() {

  rootDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
  ca := NewCA(rootDir)
  fmt.Println(ca.RootDir)
  sn, _ := ca.GetSerialNumber()
  fmt.Println(sn.Int64())
  ca.IncreaseSerialNumber()
}
