package dto

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Value  *big.Int
	Status string
}

type LogApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Status  string
}

type LogLists struct {
	TransferLogs []LogTransfer
	ApprovalLogs []LogApproval
}
