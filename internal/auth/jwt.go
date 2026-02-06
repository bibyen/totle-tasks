package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// contextKey is a private type for context keys to prevent other packages
// from accidentally overwriting context values.
type contextKey string

const (
	userIDKey contextKey = "user_id"
)

// TokenClaims stores the data we can bake into the JWT.
// Add more fields as necessary.
type TokenClaims struct {
	UserID               string `json:"user_id"`
	jwt.RegisteredClaims        // Standard claims. See https://pkg.go.dev/github.com/golang-jwt/jwt/v5#RegisteredClaims
}

// IssueAccessToken creates a new signed JWT string for a specific User ID.
func IssueAccessToken(userID string, secret []byte, duration time.Duration) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// Standard claims help with interoperability and security
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "totle-tasks-api",
		},
	}

	// Create the token using the HS256 algorithm
	// TODO: COnsider exploring alternative algorithms for better security
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	return token.SignedString(secret)
}

// ParseAndVerifyToken checks the signature and expiration of a raw string
// and returns the claims if the token is valid.
func ParseAndVerifyToken(rawToken string, secret []byte) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what we expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token is invalid or malformed")
	}

	return claims, nil
}

// GetIdentityFromContext retrieves the UserID from a context.
// It returns a boolean 'ok' so the caller can handle unauthenticated requests.
func GetIdentityFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

// InjectIdentityIntoContext takes a UserID and places it into the context.
// This is used by our interceptor.
func InjectIdentityIntoContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}
