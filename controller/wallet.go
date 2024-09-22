package controller

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/deploy"
	"example.com/go-fiber/dto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func GetBalance(c *fiber.Ctx) error {
	address := c.Params("address")

	balanceWei, err := GetBalanceAtAddress(address)
	if err != nil {
		fmt.Printf("Error to get Account 1 Balance %v", err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	balanceEther := config.ConvertWeiToEther(balanceWei)

	fmt.Println("Address : ", address)
	fmt.Println("Balance: ", balanceEther)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched balance",
		"balance": balanceEther,
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return nil

}
func GetBalanceAtAddress(address string) (*big.Int, error) {

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		log.Fatalf("Error to create ether client %v", err)
	}
	defer client.Close()
	balanceWei, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	return balanceWei, err
}

func SendBalanceToAddress(c *fiber.Ctx) error {
	type RequestBody struct {
		RecipientAddr string  `json:"recipientAddr"`
		Amount        float64 `json:"amount"`
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	sessionKey := config.GetSessionValue(c, "sessionKey")
	storedValue := config.GetSessionValue(c, "encryptedPrivateKey")
	if storedValue == nil || sessionKey == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Session value for 'encryptedPrivateKey' is nil or session key is nil",
		})
	}
	//decypt the encrypted privateKey
	encryptedPrivateKey, err := config.DecryptData(storedValue.(string), sessionKey.(string))
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error getting decrypted private key",
		})
	}
	fmt.Printf("STored key %v", encryptedPrivateKey)
	privateKey, err := config.JSONToKey(encryptedPrivateKey)
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error getting private keys from encrypted data",
		})
	}

	tx, err := SendBalance(*privateKey, body.RecipientAddr, body.Amount)

	// tx, err := SendBalanceFromGanacheAddress(body.RecipientAddr, body.Amount)

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Transaction failed: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"success":     true,
		"message":     "Transaction successfully completed",
		"transaction": tx,
	})

}
func SendBalance(key keystore.Key, recipientAddr string, amount float64) (string, error) {
	recipientAddress := common.HexToAddress(recipientAddr)
	//	publicKey := privateKey.PrivateKey.PublicKey
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	publicKey := key.PrivateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	//Check blance
	balanceWei, err := GetBalanceAtAddress(fromAddress.String())
	if err != nil {
		// log.Fatalf("Error to get Sender Balance %v", err)
		return "", err
	}

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return "", err
	}
	defer client.Close()

	gasLimit := uint64(60000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())

	fmt.Printf("GAS PRICE : %v\n", gasPrice)
	if err != nil {
		// log.Fatal("Error to get the Gas Price suggestion ", err)
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal("Error to get the Nonce from the address ", err)
		return "", err
	}
	value, err := config.ConvertEtherToWei(amount)
	if err != nil {
		// log.Fatal("Error to convert ether to wei ", err)
		return "", err
	}

	// Estimate total cost (amount + gas fees)
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	totalCost := new(big.Int).Add(value, gasCost)
	fmt.Printf("VALUE : %v\n", value)
	fmt.Printf("GAS COST : %v\n", gasCost)
	fmt.Printf("total cost : %v\n", totalCost)
	//Just for testing purposes
	privateKey := key.PrivateKey
	privateKeyBytes := privateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Printf("Private Key : %v\n", privateKeyHex)

	// Check if balance is sufficient
	if balanceWei.Cmp(totalCost) < 0 {
		return "", fmt.Errorf("insufficient balance: need %s wei, but have %s wei", totalCost.String(), balanceWei.String())
	}

	// Create a new transaction
	tx := types.NewTransaction(
		nonce, // nonce
		recipientAddress,
		value,
		gasLimit, // gas limit
		gasPrice, // gas price
		nil,      // data
	)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	//signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, key.PrivateKey)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}

func GetSessionPrivateKey(c *fiber.Ctx) (*keystore.Key, error) {
	sessionKey := config.GetSessionValue(c, "sessionKey")
	storedValue := config.GetSessionValue(c, "encryptedPrivateKey")
	if storedValue == nil || sessionKey == nil {
		return nil, errors.New("session value for encryptedPrivateKey is nil or session key is nil")
	}
	//decypt the encrypted privateKey
	encryptedPrivateKey, err := config.DecryptData(storedValue.(string), sessionKey.(string))
	if err != nil {
		return nil, errors.New("error getting decrypted private key")

	}

	privateKey, err := config.JSONToKey(encryptedPrivateKey)
	if err != nil {
		return nil, errors.New("error getting private keys from encrypted data")
	}
	return privateKey, err
}
func DeployEngageCoin(c *fiber.Ctx) error {

	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		//log.Fatalf("Error to create ether client %v", err)
		return err
	}
	defer client.Close()
	// privateKey, err := GetSessionPrivateKey(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// }

	// nonce, err := client.PendingNonceAt(context.Background(), privateKey.Address)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "Error getting Pending nonce",
	// 	})
	// }
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "Failed to get chain/network id",
	// 	})
	// }
	// fmt.Printf("CHAIN ID: %v\n", chainID)
	// fmt.Printf("NONCE : %v\n", nonce)
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "failed to get suggested gas price",
	// 	})
	// }
	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey.PrivateKey, big.NewInt(1337))
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "Failed to create authorized transactor",
	// 	})
	// }

	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)      // in wei
	// auth.GasLimit = uint64(6000000) // in units
	// auth.GasPrice = gasPrice

	auth, err := GetAuth(c, client, true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to create authorized transactor",
		})
	}
	originalOwner := auth.From

	fmt.Printf("OWNER ADDRESS : %v\n", auth.Signer)
	address, tx, instance, err := deploy.DeployEngageCoin(auth, client, originalOwner)
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

		tokenName, err := instance.Name(&bind.CallOpts{})
		if err != nil {
			log.Printf("Failed to retrieve token name: %v", err)
		}

		tokenSymbol, err := instance.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Printf("Failed to retrieve token symbol: %v", err)
		}

		totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
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

func SaveTokenInfo(tokenName string, tokenSymbol string, totalSupply *big.Int, address common.Address, userId int) error {
	query := `
	INSERT INTO tokens (token_name, symbol, total_supply,contract_address, owner_id, created_at)
	VALUES ($1, $2, $3, $4,$5, NOW())
	RETURNING id
`
	var tokenID int
	err := database.DB.QueryRow(query, tokenName, tokenSymbol, totalSupply.String(), address.Hex(), userId).Scan(&tokenID)
	if err != nil {
		log.Fatalf("Failed to insert token details into database: %v", err)
		return err
	}

	fmt.Printf("Token stored in database with ID: %d\n", tokenID)
	return nil

}

func TransferReward(c *fiber.Ctx) error {
	type RequestBody struct {
		RecipientAddr string  `json:"recipientAddr"`
		Amount        float64 `json:"amount"`
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
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	token, err := GetEngageCContractData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// key, err := GetSessionPrivateKey(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// }

	// nonce, err := client.PendingNonceAt(context.Background(), key.Address)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "failed to get suggested gas price",
	// 	})
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "Failed to create authorized transactor",
	// 	})
	// }

	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)      // in wei
	// auth.GasLimit = uint64(9000000) // in units
	// auth.GasPrice = gasPrice
	auth, err := GetAuth(c, client, true)
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
	fmt.Println("Recipient address: ", body.RecipientAddr)
	fmt.Println("Amount :", amountInt)
	tx, err := RewardUserWithTokens(client, common.HexToAddress(token.ContractAddress), auth, common.HexToAddress(body.RecipientAddr), amountInt)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("failed to transfer reward amount: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":     true,
		"message":     "successfully transfered reward amount",
		"transaction": tx,
	})
}
func RewardUserWithTokens(client *ethclient.Client, contractAddress common.Address, adminAuth *bind.TransactOpts, userAddress common.Address, tokenAmount *big.Int) (common.Hash, error) {
	// Load the token contract
	tokenInstance, err := deploy.NewEngageCoin(contractAddress, client)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to load token contract: %v", err)
	}

	fmt.Println("User Address :", userAddress)
	// Call the transfer method of the ERC-20 contract
	tx, err := tokenInstance.Transfer(adminAuth, userAddress, tokenAmount)
	// tx, err := tokenInstance.TransferFrom(adminAuth, adminAuth.From, userAddress, tokenAmount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to transfer tokens: %v", err)
	}

	fmt.Printf("Tokens transferred! Transaction hash: %s\n", tx.Hash().Hex())
	return tx.Hash(), nil
}

func GetRecentTransactions(c *fiber.Ctx) error {
	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	blogRange := uint64(1000)
	start, end, err := getStartAndEndBlock(client, int64(blogRange))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get start and end blocks : %v", err),
		})
	}

	token, err := GetEngageCContractData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get ENGC Token data from db: %v", err),
		})
	}

	fmt.Print("Token: ", token.ContractAddress)
	blocks, err := getTransactionsBetweenBlocks(client, token.ContractAddress, start, end)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to get transactions between blocks: %v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":      true,
		"message":      "successfully get the trancations",
		"transactions": blocks,
	})
}
func getTransactionsBetweenBlocks(client *ethclient.Client, tokenAddress string, startBlock, endBlock *big.Int) (dto.LogLists, error) {
	if startBlock.Cmp(big.NewInt(0)) < 0 {
		startBlock.SetInt64(0) // Cap startBlock at 0 (genesis block)
	}

	// transferEventSignature := []common.Hash{
	// 	common.HexToHash("0xddf252ad1be2c89b69c2b0691e1c1c5f6a3b3d0b17b448a7fdbf8e4c5e1b6d29"),
	// }

	// var transferLogs []dto.LogTransfer
	// var approvalLogs []dto.LogApproval
	var logLists dto.LogLists
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{common.HexToAddress(tokenAddress)},
		// Topics:    [][]common.Hash{transferEventSignature},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return logLists, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(deploy.EngageCoinABI)))
	if err != nil {
		return logLists, err
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		var transferEvent dto.LogTransfer
		approvalEvent := new(dto.LogApproval)

		txHash := vLog.TxHash

		// Get the transaction receipt
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Fatalf("Failed to retrieve transaction receipt: %v", err)
		}

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():

			// Check transaction status
			if receipt.Status == 1 {
				transferEvent.Status = "completed"

			} else if receipt.Status == 0 {
				transferEvent.Status = "failed"

			}

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			logLists.TransferLogs = append(logLists.TransferLogs, transferEvent)
		case logApprovalSigHash.Hex():

			if receipt.Status == 1 {
				approvalEvent.Status = "completed"

			} else if receipt.Status == 0 {
				approvalEvent.Status = "failed"

			}

			// approvalEvent, err := contractAbi.Unpack("Approval", vLog.Data)
			err := contractAbi.UnpackIntoInterface(approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.Owner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
			logLists.ApprovalLogs = append(logLists.ApprovalLogs, *approvalEvent)

		}
	}
	return logLists, err
}
func getStartAndEndBlock(client *ethclient.Client, blockRange int64) (*big.Int, *big.Int, error) {
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get the latest block: %v", err)
	}

	endBlock := latestBlock.Number()

	// Calculate the start block based on the block range (e.g., last 1000 blocks)
	startBlock := new(big.Int).Sub(endBlock, big.NewInt(blockRange))

	return startBlock, endBlock, nil
}

// NOTE: this is for testing on Ganache
// func senderPrivateKey() (*ecdsa.PrivateKey, error) {
// 	sender, err := crypto.HexToECDSA("fe942341ab39950bd67f208af7e22b31f51720799b95a631736eac5fd98bd378")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sender, nil
// }
// func SendBalanceFromGanacheAddress(recipientAddr string, amount float64) (string, error) {
// 	recipientAddress := common.HexToAddress(recipientAddr)
// 	privateKey, err := senderPrivateKey()
// 	if err != nil {
// 		return "", err
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return "", errors.New("failed to convert public key to ecdsa public key")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	balanceWei, err := GetBalanceAtAddress(fromAddress.String())
// 	if err != nil {
// 		return "", err
// 	}

// 	client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

// 	if err != nil {
// 		//log.Fatalf("Error to create ether client %v", err)
// 		return "", err
// 	}
// 	defer client.Close()

// 	gasLimit := uint64(30000) // in units
// 	gasPrice, err := client.SuggestGasPrice(context.Background())

// 	fmt.Printf("GAS PRICE : %v\n", gasPrice)
// 	if err != nil {
// 		// log.Fatal("Error to get the Gas Price suggestion ", err)
// 		return "", err
// 	}
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		//log.Fatal("Error to get the Nonce from the address ", err)
// 		return "", err
// 	}
// 	value, err := config.ConvertEtherToWei(amount)
// 	if err != nil {
// 		// log.Fatal("Error to convert ether to wei ", err)
// 		return "", err
// 	}

// 	// Estimate total cost (amount + gas fees)
// 	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
// 	totalCost := new(big.Int).Add(value, gasCost)
// 	fmt.Printf("VALUE : %v\n", value)
// 	fmt.Printf("GAS COST : %v\n", gasCost)
// 	fmt.Printf("total cost : %v\n", totalCost)

// 	// Check if balance is sufficient
// 	if balanceWei.Cmp(totalCost) < 0 {
// 		return "", fmt.Errorf("insufficient balance: need %s wei, but have %s wei", totalCost.String(), balanceWei.String())
// 	}

// 	// Create a new transaction
// 	tx := types.NewTransaction(
// 		nonce, // nonce
// 		recipientAddress,
// 		value,
// 		gasLimit, // gas limit
// 		gasPrice, // gas price
// 		nil,      // data
// 	)

// 	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
// 	if err != nil {
// 		// log.Fatal(err)
// 		return "", err
// 	}

// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		// log.Fatal(err)
// 		return "", err
// 	}

// 	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
// 	return signedTx.Hash().Hex(), nil
// }
