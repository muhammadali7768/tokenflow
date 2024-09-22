package controller

import (
	"database/sql"
	"fmt"
	"math/big"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/deploy"
	"example.com/go-fiber/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func GetEngageCBalance(c *fiber.Ctx) error {
	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))
	address := c.Params("address")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to connect to chain",
		})
	}
	defer client.Close()

	token, err := GetEngageCContractData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC Token data from db: %v", err),
		})
	}

	//contractAddress := common.HexToAddress("0xB0D17f2661618Bf9ABB8993daF7b7d8B5F1F1166")
	instance, err := deploy.NewEngageCoin(common.HexToAddress(token.ContractAddress), client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to get new instance",
		})
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC balance: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"balance": config.ConvertWeiToEther(balance),
		"message": "succefully get the balance",
	})
}

func GetEngageCContractData() (model.Token, error) {
	var token model.Token
	var TotalSupply []byte
	row := database.DB.QueryRow("SELECT id, token_name, symbol, total_supply, decimals, owner_id, contract_address from tokens WHERE symbol=$1", "ENGC")
	err := row.Scan(&token.ID, &token.TokenName, &token.Symbol, &TotalSupply, &token.Decimals, &token.OwnerID, &token.ContractAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return token, fmt.Errorf("no token found with symbol ENGC ")
		}
		return token, fmt.Errorf("failed to get token data: %v", err)
	}

	token.TotalSupply = new(big.Int) //first initialze to not get the invalid memory or nil pointer error
	token.TotalSupply.SetBytes(TotalSupply)
	if token.TotalSupply == nil || token.TotalSupply.Cmp(big.NewInt(0)) <= 0 {
		return token, fmt.Errorf("total supply is invalid: zero or negative value")
	}

	return token, nil
}

func GetEngcTokenDistribution(c *fiber.Ctx) error {
	type Distribution struct {
		TotalSupply *big.Float
		OwnerWallet *big.Float
		UserWallet  *big.Float
		TotalStaked *big.Float
	}
	var distribution = Distribution{
		TotalSupply: big.NewFloat(0),
		OwnerWallet: big.NewFloat(0),
		UserWallet:  big.NewFloat(0),
		TotalStaked: big.NewFloat(0),
	}
	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to connect to chain",
		})
	}
	defer client.Close()

	stakeData, err := GetTotalStackedAmount(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to stack amount %v", err),
		})
	}
	stakeCon := config.ConvertWeiToEther(stakeData)
	distribution.TotalStaked = stakeCon
	token, err := GetEngageCContractData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC Token data from db: %v", err),
		})
	}

	// Get the token contract instance
	tokenAddress := common.HexToAddress(token.ContractAddress)
	tokenInstance, err := deploy.NewEngageCoin(tokenAddress, client)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to start an ENGC token instance: %v", err),
		})
	}

	// Define the addresses to include in the token distribution
	addresses, ownerAddress, err := GetWalletAddresses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get wallet addresses: %v", err),
		})
	}

	//Get owner balance
	ownerBalance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, ownerAddress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC owner balance: %v", err),
		})
	}
	distribution.OwnerWallet = config.ConvertWeiToEther(ownerBalance)

	// Get the token balances for each address
	balances := make(map[common.Address]*big.Int)

	for _, address := range addresses {
		balance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, address)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"message": fmt.Sprintf("Failed to get ENGC Token balance: %v", err),
			})
		}
		balanceCon := config.ConvertWeiToEther(balance)
		fmt.Print("USER BALANCE :", balanceCon)
		distribution.UserWallet.Add(distribution.UserWallet, balanceCon)
		balances[address] = balance
	}

	// Calculate the total supply
	totalSupply, err := tokenInstance.TotalSupply(&bind.CallOpts{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC Token total supply: %v", err),
		})

	}

	distribution.TotalSupply = config.ConvertWeiToEther(totalSupply)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":      true,
		"message":      "successfully get the disctibution",
		"distribution": distribution,
	})
}

func GetWalletAddresses() ([]common.Address, common.Address, error) {

	var addresses []common.Address
	var ownerAddress common.Address
	rows, err := database.DB.Query("SELECT ethereum_address, role from users")
	if err != nil {
		if err == sql.ErrNoRows {
			return addresses, ownerAddress, fmt.Errorf("no token found with symbol ENGC ")
		}
		return addresses, ownerAddress, fmt.Errorf("failed to get token data: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var address string
		var role string
		rows.Scan(&address, &role)
		if role == "owner" {
			ownerAddress = common.HexToAddress(address)
		} else {
			addresses = append(addresses, common.HexToAddress(address))
		}
	}

	return addresses, ownerAddress, nil
}
