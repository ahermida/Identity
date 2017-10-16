package main

import (
  "bytes"
  "crypto/elliptic"
  "encoding/gob"
  "fmt"
  "io/ioutil"
  "log"
  "os"
)

const walletFile = "wallet_%s.dat"

type Wallets struct {
  Wallets map[string]*Wallet
}

func NewWallets(nodeID string) (*Wallets, error) {
  wallets := Wallets{}
  wallets.Wallets = make(map[string]*Wallet)

  err := wallets.LoadFromFile(nodeID)

  return &wallets, err
}

// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
  var addresses []string

  for address := range ws.Wallets {
    addresses = append(addressesm address)
  }

  return addresses
}

// Loads wallets from a file
func (ws *Wallets) LoadFromFile(nodeID string) error {
  walletFile := fmt.Sprintf(walletFile, nodeID)
  if _, err := os.Stat(walletFile); os.IsNotExist(err) {
    return err
  }

  fileContent, err := ioutil.ReadFile(walletFile)
  if err != nil {
    log.Panic(err)
  }

  var wallets Wallets
  gob.Register(elliptic.P256())
  decoder := gob.NewDecoder(bytes.NewReader(fileContent))
  err = decoder.Decode(&wallets)
  if err != nil {
    log.Panic(err)
  }

  ws.Wallets = wallets.Wallets

  return nil
}

// Saves wallets to a file
func (ws Wallets) SaveToFile(nodeID string) {
  var content bytes.Buffer
  walletFile := fmt.Sprintf(walletFile, nodeID)

  gob.Register(elliptic.P256())

  encoder := gob.NewEncoder(&content)
  err := encoder.Encode(ws)
  if err != nil {
    log.Panic(err)
  }

  err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
  if err != nil {
    log.Panic(err)
  }
}
