package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/deploy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func DeployENGCStaking(c *fiber.Ctx) error {

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return err
	}
	defer client.Close()

	rewardRate := big.NewInt(10000000)
	auth, err := GetAuth(c, client, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to create authorized transactor",
		})
	}

	address, tx, instance, err := deploy.DeployENGCStaking(auth, client, common.HexToAddress("0x2106b46bE3f0D4ddFf1bb9b4AF5eB2c1036E8F35"), rewardRate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	fmt.Printf("TRANSACTION: %v\n", tx)
	fmt.Printf("Address: %v\n", address)

	response := fiber.Map{
		"success":          true,
		"message":          "Contract deployment initiated",
		"transaction_hash": tx.Hash().Hex(),
		"contract_address": address.Hex(),
	}

	userId, err := GetUserId(c)
	if err != nil {
		fmt.Print("USER ID ERROR", err)
	}

	go func() {
		fmt.Println("Waiting for the contract deployment to be mined...")
		receipt, err := bind.WaitMined(context.Background(), client, tx)

		if err != nil {
			log.Printf("Transaction failed to be mined: %v", err)
			// Handle error or notify client of failure
		}

		if receipt.Status == 1 {
			log.Printf("Transaction mined successfully! Contract Address: %s", address.Hex())
			// TODO: Notify the client through websockets
		} else {
			log.Printf("Transaction reverted. Contract Address: %s", address.Hex())
		}

		if err != nil {
			log.Fatalf("Failed to load contract instance: %v", err)
		}

		tokenName := "ENGC Stacking"

		tokenSymbol := "No symbol"

		totalSupply, err := instance.TotalStaked(&bind.CallOpts{})
		if err != nil {
			log.Printf("Failed to retrieve total supply: %v", err)
		}

		err = SaveTokenInfo(tokenName, tokenSymbol, totalSupply, address, userId)
		if err != nil {
			log.Printf("Failed to retrieve total supply: %v", err)
		}
	}()

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAuth(c *fiber.Ctx, client *ethclient.Client, allOptions bool) (*bind.TransactOpts, error) {
	privateKey, err := GetSessionPrivateKey(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get the private key from session %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	fmt.Print("Chain ID", chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get the chainid %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey.PrivateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the transaction with private key %v", err)
	}

	if allOptions {
		nonce, err := client.PendingNonceAt(context.Background(), privateKey.Address)
		if err != nil {
			return nil, fmt.Errorf("failed to get the nounce value %v", err)
		}
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get the gas price %v", err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)       // in wei
		auth.GasLimit = uint64(10000000) // in units
		auth.GasPrice = gasPrice
	}
	return auth, nil
}
func DepositAmount(c *fiber.Ctx) error {
	type RequestBody struct {
		Amount float64 `json:"amount"`
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Invalid data %v", err),
		})
	}

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return err
	}
	defer client.Close()

	stackAddress, err := GetStackContractAddresses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to get the stack address %v", err),
		})
	}
	stackingInstance, err := deploy.NewENGCStaking(stackAddress, client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to start the NewEngStack instance %v", err),
		})
	}

	engcTokenCont, err := GetEngageCContractData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to get the ENGC contract address %v", err),
		})
	}

	engcToken, err := deploy.NewEngageCoin(common.HexToAddress(engcTokenCont.ContractAddress), client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to start the ENGC instance %v", err),
		})
	}

	auth, err := GetAuth(c, client, true)
	fmt.Print("AUTH", auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to create authorized transactor",
		})
	}
	amountInt, err := config.ConvertEtherToWei(body.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "failed to convert amount to big int",
		})
	}
	// Approve staking contract to spend tokens
	tx, err := engcToken.Approve(auth, stackAddress, amountInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to approve staking contract to spend tokens %v", err),
		})
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("failed to wait for approval transaction to be mined: %v", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("approval transaction failed: %v", receipt.Status)
	}

	fmt.Println("Approval transaction mined successfully")
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("failed to wait for approval transaction to be mined: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	dtx, err := stackingInstance.Deposit(auth, amountInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to deposit the amount to stacking address %v", err),
		})
	}

	fmt.Println("Amount approved for stacking ", tx)
	fmt.Println("Amount deposited for stacking ", dtx)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":          true,
		"message":          fmt.Sprintf("successfully deposited the amount %v", err),
		"transaction_hash": dtx,
	})
}

func GetStakedAmount(c *fiber.Ctx) error {
	address := c.Params("address")

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return err
	}
	defer client.Close()

	stackAddress, err := GetStackContractAddresses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to get the stack address %v", err),
		})
	}
	stackingInstance, err := deploy.NewENGCStaking(stackAddress, client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to start the NewEngStack instance %v", err),
		})
	}

	stakeData, err := stackingInstance.Stakes(&bind.CallOpts{}, common.HexToAddress(address))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to stack amount %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "successfully get the stake data",
		"stake":   stakeData,
	})
}

func GetTotalStacked(c *fiber.Ctx) error {

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return err
	}
	defer client.Close()

	stakeData, err := GetTotalStackedAmount(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to stack amount %v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "successfully get the total stake",
		"stake":   stakeData,
	})
}
func GetTotalStackedAmount(client *ethclient.Client) (*big.Int, error) {
	stackAddress, err := GetStackContractAddresses()
	if err != nil {
		return nil, fmt.Errorf("failed to get the stack address %v", err)
	}
	stackingInstance, err := deploy.NewENGCStaking(stackAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to start the NewEngStack instance %v", err)
	}
	stakeData, err := stackingInstance.TotalStaked(&bind.CallOpts{})
	return stakeData, err
}

func GetStackContractAddresses() (common.Address, error) {

	var ownerAddress string
	row := database.DB.QueryRow("SELECT contract_address from tokens where token_name=$1", "ENGC Stacking")
	err := row.Scan(&ownerAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return common.HexToAddress(ownerAddress), fmt.Errorf("no stacking address found")
		}
		return common.HexToAddress(ownerAddress), fmt.Errorf("failed to get token data: %v", err)
	}

	return common.HexToAddress(ownerAddress), nil
}
