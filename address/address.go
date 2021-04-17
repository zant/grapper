package address

import (
  "context"
  "math/big"

  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/zant/geth-wrapper/utils"
)

type Address struct {
  client  *ethclient.Client
  address common.Address
}

func NewAddress(client *ethclient.Client, addStr string) Address {
  address := common.HexToAddress(addStr)
  return Address{client, address}
}

func (a *Address) Balance(block *big.Int) (*big.Float, error) {
  balance, err := a.client.BalanceAt(context.Background(), a.address, block)
  return utils.WeiToEth(balance), err
}

func (a *Address) PendingBalance() (*big.Float, error) {
  balance, err := a.client.PendingBalanceAt(context.Background(), a.address)
  return utils.WeiToEth(balance), err
}

func (a *Address) IsContract() (bool, error) {
  bytecode, err := a.client.CodeAt(context.Background(), a.address, nil)
  return len(bytecode) > 0, err
}
