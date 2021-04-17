package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/zant/grapper/wallet"
)

func main() {
  client, err := ethclient.Dial("https://mainnet.infura.io/v3/d851d38fbced4fb2bd8ea0d049af97a1")
  if err != nil {
    log.Fatal(err)
  }
  _ = client

  txWallet, err := wallet.NewWalletFromFile(".wallets/pk1")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(txWallet.Address())
}
