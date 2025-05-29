package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Order represents an order on the ClusterMarket contract.
type Order struct {
	Id             *big.Int
	User           common.Address
	Status         uint8
	Ip             *big.Int
	Gpu            *big.Int
	Cpu            *big.Int
	MemoryBytes    *big.Int
	Disk           *big.Int
	Network        *big.Int
	RentalDuration *big.Int
	PaymentToken   common.Address
	ClusterId      *big.Int
	PaidAmount     *big.Int
	DiscountAmount *big.Int
	OrderType      uint8
}
