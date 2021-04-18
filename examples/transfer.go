package main

import (
  "fmt"
  "log"
  "math/big"

  "github.com/zant/grapper/wallet"
)

func main() {
  // Read from saved pk & connecto to ganache
  w, err := wallet.NewWalletFromFile("HTTP://127.0.0.1:7545", ".wallets/pk1")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(w.Balance())
  // Transfer to address
  value := big.NewInt(1000000000000000000)
  err = w.Transfer("0x6E45c47bd6Dc099EBdbd95C270323747b55FEC09", value)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(w.Balance())
}
