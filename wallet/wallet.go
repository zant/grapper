package wallet

import (
  "crypto/ecdsa"
  "log"

  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
  privateKey *ecdsa.PrivateKey
  publicKey  *ecdsa.PublicKey
}

func PublicKeyFromPrivate(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
  publicKeyCrypto := privateKey.Public()
  publicKey, ok := publicKeyCrypto.(*ecdsa.PublicKey)
  if !ok {
    log.Fatal("Cannot assert public key")
  }
  return publicKey
}

func NewWallet() Wallet {
  privateKey, err := crypto.GenerateKey()
  if err != nil {
    log.Fatal(err)
  }
  publicKey := PublicKeyFromPrivate(privateKey)
  return Wallet{privateKey, publicKey}
}

func NewWalletFromFile(file string) (Wallet, error) {
  privateKey, err := crypto.LoadECDSA(file)
  publicKey := PublicKeyFromPrivate(privateKey)
  return Wallet{privateKey, publicKey}, err
}

func (w *Wallet) PrivateKeyHex() string {
  privateKeyBytes := crypto.FromECDSA(w.privateKey)
  return hexutil.Encode(privateKeyBytes[2:])
}

func (w *Wallet) PublicKeyHex() string {
  privateKeyBytes := crypto.FromECDSAPub(w.publicKey)
  return hexutil.Encode(privateKeyBytes[2:])
}

func (w *Wallet) Address() string {
  return crypto.PubkeyToAddress(*w.publicKey).Hex()
}

func (w *Wallet) Save(file string) error {
  return crypto.SaveECDSA(file, w.privateKey)
}
