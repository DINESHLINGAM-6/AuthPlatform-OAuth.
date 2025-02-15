package handlers

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// Define a secret key for signing the JWT (keep this secure in production)
var jwtKey = []byte("your-secret-key")

// Claims struct to encode/decode the JWT payload
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for the given email
func GenerateJWT(email string) (string, error) {
	// Set token expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-system",
		},
	}

	// Create the token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}