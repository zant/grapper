package wallet

import (
  "context"
  "crypto/ecdsa"
  "fmt"
  "log"
  "math/big"

  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/zant/grapper/utils"
)

type Wallet struct {
  client     *ethclient.Client
  address    common.Address
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

func NewWallet(client *ethclient.Client) (Wallet, error) {
  privateKey, err := crypto.GenerateKey()
  publicKey := PublicKeyFromPrivate(privateKey)
  address := crypto.PubkeyToAddress(*publicKey)
  return Wallet{client, address, privateKey, publicKey}, err
}

func NewWalletFromFile(rpcServer string, file string) (Wallet, error) {
  client, err := ethclient.Dial(rpcServer)
  if err != nil {
    log.Fatal(err)
  }
  privateKey, err := crypto.LoadECDSA(file)
  publicKey := PublicKeyFromPrivate(privateKey)
  address := crypto.PubkeyToAddress(*publicKey)
  return Wallet{client, address, privateKey, publicKey}, err
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

func (w *Wallet) Balance() (*big.Float, error) {
  balance, err := w.client.BalanceAt(context.Background(), w.address, nil)
  return utils.WeiToEth(balance), err
}

func (w *Wallet) PendingBalance() (*big.Float, error) {
  balance, err := w.client.PendingBalanceAt(context.Background(), w.address)
  return utils.WeiToEth(balance), err
}

func (w *Wallet) IsContract() (bool, error) {
  bytecode, err := w.client.CodeAt(context.Background(), w.address, nil)
  return len(bytecode) > 0, err
}

func (w *Wallet) Transfer(address string, value *big.Int) error {
  fromAddress := crypto.PubkeyToAddress(*w.publicKey)
  toAddress := common.HexToAddress(address)
  gasLimit := uint64(21000)

  nonce, err := w.client.PendingNonceAt(context.Background(), fromAddress)
  if err != nil {
    log.Fatal(err)
  }
  gasPrice, err := w.client.SuggestGasPrice(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
  chainId, err := w.client.NetworkID(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), w.privateKey)
  if err != nil {
    log.Fatal(err)
  }

  err = w.client.SendTransaction(context.Background(), signedTx)

  fmt.Printf("Tx sent: %s", signedTx.Hash().Hex())
  return err
}
