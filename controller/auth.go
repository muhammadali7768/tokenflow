package controller

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/dto"
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

	row := database.DB.QueryRow("SELECT id, email, username, password_hash, ethereum_address, role FROM users WHERE email=$1", input.Email)
	row.Scan(&storedUser.ID, &storedUser.Email, &storedUser.Username, &storedUser.Password, &storedUser.Wallet, &storedUser.Role)

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "Invalid credentials",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": storedUser.ID,
		"name":    storedUser.Username,
		"role":    storedUser.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
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

	dtoUser := dto.UserResponse{
		ID:       storedUser.ID,
		Email:    storedUser.Email,
		Username: storedUser.Username,
		Wallet:   storedUser.Wallet,
		Role:     storedUser.Role,
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		Secure:   true,
	})
	fmt.Println("Setting session key key code is fine")
	// return c.JSON(fiber.Map{"success": true, "token": t, "user": dtoUser})
	return c.JSON(fiber.Map{"success": true, "user": dtoUser})
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
        INSERT INTO users (username, email, password_hash, ethereum_address, keystore_file, encrypted_mnemonic, role)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err = database.DB.Exec(query, username, email, passwordHash, ethereumAddress, keystoreFile, encryptedMnemonic, "user")
	if err != nil {
		return err
	}

	return nil
}

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

func GetUserId(c *fiber.Ctx) (int, error) {
	// authHeader := c.Get("Authorization")

	// fmt.Print(authHeader)
	// if authHeader == "" {
	// 	return 0, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
	// }

	// // Extract token string (remove "Bearer " prefix if present)
	// tokenString := authHeader
	// if strings.HasPrefix(authHeader, "Bearer ") {
	// 	tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	// }
	token := c.Cookies("token")

	fmt.Println("Token", token)
	// Parse token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Config("JWT_SECRET")), nil
	})

	fmt.Println("Parsed token", parsedToken)
	if err != nil || !parsedToken.Valid {
		return 0, err
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Printf("Clains %v\n, %v\n", parsedToken, ok)
	if !ok {
		return 0, err
	}

	// Extract user ID
	userID, ok := claims["user_id"].(float64) // JWT claims are typically in float64 format
	fmt.Println("USER ID AFTER CLEANING ", userID, ok)
	if !ok {
		return 0, errors.New("user id missing from token")
	}
	return int(userID), nil
}
