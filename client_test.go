package main

import (
	"math/big"
	"testing"

	"github.com/velas/Velas-Sphere-Client/types"

	"github.com/stretchr/testify/assert"
)

func TestClient_Deposit(t *testing.T) {
	t.Run("invalid key", func(t *testing.T) {
		key := "0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
		contractAddr := "0x99665E952674a8ba03e8e91A09B0163be5DcfB5A"
		server := "http://localhost:8545"
		_, err := NewClient(key, contractAddr, server)
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to parse private key: invalid hex string")
	})

	t.Run("valid key", func(t *testing.T) {
		key := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
		contractAddr := "0x99665E952674a8ba03e8e91A09B0163be5DcfB5A"
		server := "http://localhost:8545"
		client, err := NewClient(key, contractAddr, server)
		assert.NoError(t, err)
		_, err = client.Deposit(1000)
		assert.NoError(t, err)
	})
}

func TestClient_RegisterNode(t *testing.T) {
	key := "1c83e027c0540f38f298213f9da4ee6acd36d5342058d18de39bd54e3bea4d71"
	contractAddr := "0x99665E952674a8ba03e8e91A09B0163be5DcfB5A"
	server := "http://localhost:8545"
	client, err := NewClient(key, contractAddr, server)
	assert.NoError(t, err)
	pricing := types.Pricing{
		KeepPerByte:   big.NewInt(1),
		WritePerByte:  big.NewInt(1),
		GPUTPerCycle:  big.NewInt(1),
		CPUTtPerCycle: big.NewInt(1),
	}
	_, err = client.RegisterNode(pricing)
	assert.NoError(t, err)
}
