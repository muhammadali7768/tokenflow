package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tyler-smith/go-bip39"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Validate the input data
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username, email, and password are required",
		})
	}

	err := CreateNewUser(user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User registered successfully",
	})

}

func LoginUser(c *fiber.Ctx) error {
	var input model.User
	var storedUser model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	row := database.DB.QueryRow("SELECT id, username, password_hash, ethereum_address FROM users WHERE username=$1", input.Username)
	row.Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password, &storedUser.Wallet)

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": storedUser.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error signing token with secret key",
		})
	}
	keys, err := GetPrivateKey(storedUser.ID, input.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error getting wallet keys",
		})
	}
	fmt.Println("Getting Private key code is fine")

	jsonKey, err := config.KeyToJSON(keys)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error to convert Keys to Json",
		})
	}
	fmt.Println("Getting JSON key code is fine")
	sessionKey, err := config.GenerateSessionKey()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error generating session key",
		})
	}
	fmt.Println("Generating session key code is fine")
	encryptedPrivateKey, err := config.EncryptData(jsonKey, sessionKey)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error encrypting data",
		})
	}
	values := map[string]interface{}{"encryptedPrivateKey": encryptedPrivateKey, "sessionKey": sessionKey}
	fmt.Println("Encrypted Private key code is fine")
	if err := config.SetSessionValues(c, values); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Error setting encryptedPrivateKey session data",
		})
	}
	// fmt.Println("Setting session Private key code is fine")
	// if err := config.SetSessionValue(c, "sessionKey", sessionKey); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	// 		"success": false,
	// 		"message": "Error setting session key data",
	// 	})
	// }
	fmt.Println("Setting session key key code is fine")
	return c.JSON(fiber.Map{"token": t, "wallet_address": storedUser.Wallet})
}

// Function to generate a new user account with mnemonic and keystore
func CreateNewUser(username, email, password string) error {
	//  Generate entropy
	entropy, err := bip39.NewEntropy(256) // 24-word mnemonic
	if err != nil {
		return err
	}

	//  Generate mnemonic
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}

	// Encrypt the mnemonic
	encryptedMnemonic, err := config.EncryptData(mnemonic, password)
	if err != nil {
		return err
	}

	//  Derive the seed from the mnemonic
	seed := bip39.NewSeed(mnemonic, password)

	//  Use the seed to generate a private key
	privateKey, err := crypto.ToECDSA(seed[:32]) // Use the first 32 bytes for the private key
	if err != nil {
		return err
	}

	// Create an account using the private key and store it in the keystore
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return err
	}

	// 7. Store user details in the database
	ethereumAddress := account.Address.Hex()
	keystoreFile := account.URL.Path // e.g., "UTC--2021-09-01T00-00-00.000000000Z--address"

	// Hash the password before storing (use bcrypt or similar)
	passwordHash, err := getHashedPassword(password)
	if err != nil {
		return err
	}

	// Insert into PostgreSQL
	query := `
        INSERT INTO users (username, email, password_hash, ethereum_address, keystore_file, encrypted_mnemonic)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err = database.DB.Exec(query, username, email, passwordHash, ethereumAddress, keystoreFile, encryptedMnemonic)
	if err != nil {
		return err
	}

	return nil
}

// func EncryptMnemonic(mnemonic, password string) (string, error) {
// 	salt := []byte("your-salt") // Use a securely generated salt
// 	key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	aesGCM, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonce := make([]byte, aesGCM.NonceSize())
// 	if _, err := rand.Read(nonce); err != nil {
// 		return "", err
// 	}

// 	ciphertext := aesGCM.Seal(nonce, nonce, []byte(mnemonic), nil)
// 	return base64.StdEncoding.EncodeToString(ciphertext), nil
// }

// // DecryptMnemonic decrypts the mnemonic using AES.
// func DecryptMnemonic(encryptedMnemonic, password string) (string, error) {
// 	salt := []byte("your-salt") // Use the same salt used during encryption
// 	key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

// 	data, err := base64.StdEncoding.DecodeString(encryptedMnemonic)
// 	if err != nil {
// 		return "", err
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	aesGCM, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonceSize := aesGCM.NonceSize()
// 	if len(data) < nonceSize {
// 		return "", errors.New("ciphertext too short")
// 	}

// 	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
// 	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(plaintext), nil
// }

func getHashedPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashedPassword, err
}

func GetPrivateKey(userID int, password string) (*keystore.Key, error) {
	// Query the database
	var keystoreFile string
	var encryptedMnemonic string
	err := database.DB.QueryRow("SELECT keystore_file, encrypted_mnemonic FROM users WHERE id = $1", userID).Scan(&keystoreFile, &encryptedMnemonic)
	if err != nil {
		return nil, err
	}

	// Read the keystore file

	keyJSON, err := os.ReadFile(keystoreFile)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(keyJSON))

	decryptedMnemonic, _ := config.DecryptData(encryptedMnemonic, password)
	fmt.Printf("Decrypted Mnemonics %v", decryptedMnemonic)
	// Decrypt the keystore file to get the key
	key, err := keystore.DecryptKey(keyJSON, password) //TODO: we should use decryptedMnemonic for decryption in future
	if err != nil {
		return nil, err
	}
	fmt.Printf("Decrypted KEY %v", key)
	return key, nil
}
