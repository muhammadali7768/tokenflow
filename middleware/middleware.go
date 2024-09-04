package middleware

import (
	"errors"
	"fmt"
	"time"

	"example.com/go-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) error {
	fmt.Printf("User %s Pass %s", config.Config("USERNAME"), config.Config("PASSWORD"))
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}

var jwtSecret = []byte("secret")

// AuthMiddleware validates the JWT token passed in the header
var AuthMiddleware = keyauth.New(keyauth.Config{
	Validator: func(c *fiber.Ctx, token string) (bool, error) {
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !parsedToken.Valid {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}

		exp := int64(claims["exp"].(float64))
		expirationTime := time.Unix(exp, 0)
		if time.Now().After(expirationTime) {
			return false, keyauth.ErrMissingOrMalformedAPIKey
		}

		return true, nil
	},
})
