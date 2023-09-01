package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

var jwtSecretKey = "your-secret-key" // Consider using environment variables

func GenerateJWTToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TokenAuth(c *revel.Controller, actionName string) revel.Result {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return c.Forbidden("Authorization header missing")
	}

	// Extract the token from the "Bearer <token>" format
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ") //check

	// // Check if the token key exists in the cache
	// cacheKey := "token_" + tokenStr

	// // Define a placeholder value
	// var placeholderValue interface{}

	// // Use cache.Get to check if the key exists
	// if err := cache.Get(cacheKey, &placeholderValue); err != nil {
	// 	return c.Forbidden("Access Denied: Invalid or expired token")
	// }

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return c.Forbidden("Access Denied: Invalid access")
	}
	if !token.Valid {
		return c.Forbidden("Access Denied: expired token")
	}

	// Continue to the action
	return nil
}

// ... other middleware functions ...

// -vv
