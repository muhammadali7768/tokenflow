package config

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
)

var Store *session.Store

// Initialize the session store
func InitSessionStore() {
	storage := memory.New()
	Store = session.New(session.Config{
		Storage:      storage,
		Expiration:   24 * time.Hour,
		CookieSecure: true, // Ensures the cookie is only sent over HTTPS
	})
}

// Middleware to attach session store to context
func SessionMiddleware(c *fiber.Ctx) error {
	if Store == nil {
		fmt.Println("Session store not initialized")
		return fmt.Errorf("session store not initialized")
	}

	sess, err := Store.Get(c)
	if err != nil {
		fmt.Println("Error getting session:", err)
		return err
	}

	fmt.Println("Session retrieved:", sess)
	c.Locals("session", sess)
	return c.Next()
}

func SetSessionValue(c *fiber.Ctx, key string, value interface{}) error {
	sessLocal := c.Locals("session")

	// Safe type assertion to *session.Session
	sess, ok := sessLocal.(*session.Session)
	if !ok || sess == nil {
		return fmt.Errorf("session is not set or invalid type")
	}
	fmt.Printf("setting session key value, %v , %v", key, value)
	sess.Set(key, value)
	if err := sess.Save(); err != nil {
		return err
	}
	return nil
}
func SetSessionValues(c *fiber.Ctx, values map[string]interface{}) error {
	sessLocal := c.Locals("session")
	if sessLocal == nil {
		return fmt.Errorf("session is not initialized")
	}

	// Safe type assertion to *session.Session
	sess, ok := sessLocal.(*session.Session)
	if !ok || sess == nil {
		return fmt.Errorf("session is not set or invalid type")
	}

	// Iterate over the map and set each key-value pair
	for key, value := range values {
		fmt.Printf("Setting session key-value pair: %v, %v\n", key, value)
		sess.Set(key, value)
	}

	// Save the session once after all values are set
	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}

func GetSessionValue(c *fiber.Ctx, key string) interface{} {
	sessLocal := c.Locals("session")

	sess, ok := sessLocal.(*session.Session)
	if !ok || sess == nil {
		fmt.Println("Session is not set or invalid type")
		return nil
	}

	// Retrieve the value from session
	storedValue := sess.Get(key)
	if storedValue == nil {
		fmt.Printf("Session value for '%s' is nil\n", key)
		return nil
	}

	return storedValue
}

func DeleteSessionValue(c *fiber.Ctx, key string) error {
	sess := c.Locals("session").(*session.Session)
	sess.Delete(key)
	return sess.Save()
}

func ClearSession(c *fiber.Ctx) error {
	sess := c.Locals("session").(*session.Session)
	return sess.Destroy()
}
