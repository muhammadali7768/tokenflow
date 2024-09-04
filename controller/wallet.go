package controller

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"example.com/go-fiber/config"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

func GetBalance(c *fiber.Ctx) error {
	address := c.Params("address")
	// client, err := ethclient.Dial(config.Config("BLOCKCHAIN_URL"))

	// if err != nil {
	// 	log.Fatalf("Error to create ether client %v", err)
	// }
	// defer client.Close()

	// balanceWei, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)

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

	gasLimit := uint64(30000) // in units
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

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, key.PrivateKey)
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
