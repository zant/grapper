package main

import (
  "fmt"
  "log"
  "math/big"
  "time"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/zant/grapper/wallet"
)

func main() {
  w, err := wallet.NewWalletFromKeyStore("http://127.0.0.1:7545", "./.account")

  if err != nil {
    log.Fatalln(err)
  }

  nonce, err := w.Nonce()

  if err != nil {
    log.Fatalln(err)
  }

  auth, err := bind.NewKeyedTransactorWithChainID(w.PrivateKey(), big.NewInt(int64(5777)))
  auth.Nonce = big.NewInt(int64(nonce))
  auth.Value = big.NewInt(0)
  auth.GasLimit = uint64(3000000)
  auth.GasPrice = big.NewInt(1000000)

  if err != nil {
    log.Fatalln(err)
  }

  passphrase := "passphrase"
  address, tx, store, err := DeployStore(auth, w.Client(), passphrase)

  if err != nil {
    log.Fatal(err)
  }

  time.Sleep(250 * time.Millisecond)
  fmt.Println(address, tx)
  items, err := store.Items(&bind.CallOpts{}, [32]byte{})

  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item)
  }

  store.SetItem(&bind.TransactOpts{}, [32]byte{0x1}, [32]byte{0x2})

  for _, item := range items {
    fmt.Println(item)
  }
}
