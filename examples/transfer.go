package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/zant/grapper/wallet"
)

func main() {
  // Connect to ganache
  client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
  if err != nil {
    log.Fatal(err)
  }
  _ = client

  // Read from saved pk
  w, err := wallet.NewWalletFromFile(client, ".wallets/pk1")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(w.Balance())
  // Transfer to address
  err = w.Transfer("0x6E45c47bd6Dc099EBdbd95C270323747b55FEC09")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(w.Balance())
}
