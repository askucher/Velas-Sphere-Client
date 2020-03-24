package types

import "math/big"

type Pricing struct {
	KeepPerByte   *big.Int
	WritePerByte  *big.Int
	GPUTPerCycle  *big.Int
	CPUTtPerCycle *big.Int
}
