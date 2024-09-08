package model

import "math/big"

type Token struct {
	ID              int      `json:"id"`
	TokenName       string   `json:"token_name"`
	Symbol          string   `json:"symbol"`
	TotalSupply     *big.Int `json:"total_supply"`
	Decimals        int      `json:"decimals"`
	OwnerID         int      `json:"owner_id"`
	ContractAddress string   `json:"contract_address"`
}
