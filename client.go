package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/velas/Velas-Sphere-Client/types"

	contracts "github.com/velas/Velas-Sphere-Client/internal/contracts"

	"github.com/pkg/errors"

	common "github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	eth      *ethclient.Client
	key      *ecdsa.PrivateKey
	contract *contracts.Ethdepositcontract
	gasLimit uint64
	gasPrice *big.Int
}

const (
	MembershipFee   = 100000000000
	DefaultGasLimit = uint64(210000)
)

func NewClient(userKey, contractAddress, server string) (*Client, error) {
	var client Client
	eth, err := ethclient.Dial(server)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to dial %s", server))
	}
	client.eth = eth

	privateKey, err := crypto.HexToECDSA(userKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse private key")
	}

	client.key = privateKey

	if !ethereum.IsHexAddress(contractAddress) {
		return nil, errors.New("invalid contract address")
	}

	contract, err := contracts.NewEthdepositcontract(ethereum.HexToAddress(contractAddress), eth)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create contract")
	}
	client.contract = contract
	client.gasLimit = DefaultGasLimit
	return &client, nil
}

func (c *Client) WithGasPrice(gasPrice int64) *Client {
	c.gasPrice = big.NewInt(gasPrice)
	return c
}
func (c *Client) WithGasLimit(gasLimit uint64) *Client {
	c.gasLimit = gasLimit
	return c
}

// Deposit deposits amount (in gwei) to contract
func (c *Client) Deposit(amount int64) (string, error) {
	opts := c.newSignedTransactOpts()
	opts.Value = big.NewInt(amount)

	tx, err := c.contract.DepositWithNodes(opts, big.NewInt(0), big.NewInt(0))
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) DepositWithNodes(amount, pool, places int64) (string, error) {
	opts := c.newSignedTransactOpts()
	opts.Value = big.NewInt(amount)

	tx, err := c.contract.DepositWithNodes(opts, big.NewInt(pool), big.NewInt(places))
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) ProposePricing(pricing types.Pricing) (string, error) {
	opts := c.newSignedTransactOpts()

	tx, err := c.contract.ProposePricing(
		opts,
		pricing.KeepPerByte,
		pricing.WritePerByte,
		pricing.CPUTtPerCycle,
		pricing.GPUTPerCycle)
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) ChangePool(pool, places int64) (string, error) {
	opts := c.newSignedTransactOpts()

	tx, err := c.contract.ChangePool(opts, big.NewInt(pool), big.NewInt(places))
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) CreateInvoice(invoice types.Invoice) (string, error) {
	if !ethereum.IsHexAddress(invoice.User) {
		return "", errors.New("invalid user address")
	}
	user := ethereum.HexToAddress(invoice.User)

	opts := c.newSignedTransactOpts()

	tx, err := c.contract.CreateInvoice(
		opts,
		invoice.HeightStart,
		invoice.HeightEnd, user,
		invoice.Resources.KeepPerByte,
		invoice.Resources.WritePerByte,
		invoice.Resources.CPUTtPerCycle,
		invoice.Resources.GPUTPerCycle)
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) RegisterNode(pricing types.Pricing) (string, error) {
	node := crypto.PubkeyToAddress(c.key.PublicKey)
	opts := c.newSignedTransactOpts()
	opts.Value = big.NewInt(MembershipFee)

	tx, err := c.contract.RegisterNode(
		opts,
		node,
		pricing.KeepPerByte,
		pricing.WritePerByte,
		pricing.CPUTtPerCycle,
		pricing.GPUTPerCycle,
	)
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}
func (c *Client) ChangeNodePricing(pricing types.Pricing) (string, error) {
	opts := c.newSignedTransactOpts()

	tx, err := c.contract.ChangeNodePricing(
		opts,
		pricing.KeepPerByte,
		pricing.WritePerByte,
		pricing.CPUTtPerCycle,
		pricing.GPUTPerCycle,
	)
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}
func (c *Client) Withdraw(address string) (string, error) {
	if !ethereum.IsHexAddress(address) {
		return "", errors.New("invalid node address")
	}
	user := ethereum.HexToAddress(address)

	opts := c.newSignedTransactOpts()
	tx, err := c.contract.Withdraw(opts, user)
	if err != nil {
		return "", errors.Wrap(err, "transaction failed")
	}
	return tx.Hash().String(), nil
}

func (c *Client) newSignedTransactOpts() *common.TransactOpts {
	auth := common.NewKeyedTransactor(c.key)
	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	auth.Context = context.TODO()
	return auth
}
