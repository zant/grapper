package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/zant/grapper/wallet"
)

func main() {
  client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
  if err != nil {
    log.Fatal(err)
  }
  _ = client

  txWallet, err := wallet.NewWalletFromFile(client, ".wallets/pk1")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(txWallet.Address())
  fmt.Println(txWallet.Balance())
}
