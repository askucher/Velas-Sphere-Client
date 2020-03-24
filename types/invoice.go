package types

import (
	"math/big"
)

type Invoice struct {
	HeightStart *big.Int
	HeightEnd   *big.Int
	User        string
	Resources   Resources
}

type Resources struct {
	KeepPerByte   *big.Int
	WritePerByte  *big.Int
	GPUTPerCycle  *big.Int
	CPUTtPerCycle *big.Int
}
