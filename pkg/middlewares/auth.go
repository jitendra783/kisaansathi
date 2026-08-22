package middlewares

import (
	"context"
	"errors"
	"kisaanSathi/pkg/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type middlewareKey string

const (
	ContextKey middlewareKey = "claims"
)

var jwtSecret []byte

// InitJWT loads the JWT secret after config.Load() has completed.
func InitJWT() {
	jwtSecret = []byte(config.GetConfig().GetString("jwt.secret"))
}

type Claims struct {
	UserID    string `json:"userId"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"tokenType"`

	jwt.RegisteredClaims
}

type Middleware struct{}

type AuthMiddlewareInterface interface {
	AuthMiddleware(next http.Handler) http.Handler
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header missing",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetClaims(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(ContextKey).(*Claims)
	return claims, ok
}

func GetUserID(ctx context.Context) (string, bool) {
	claims, ok := GetClaims(ctx)
	if !ok {
		return "", false
	}
	return claims.UserID, true
}

func GetEmail(ctx context.Context) (string, bool) {
	claims, ok := GetClaims(ctx)
	if !ok {
		return "", false
	}
	return claims.Email, true
}

func GetRole(ctx context.Context) (string, bool) {
	claims, ok := GetClaims(ctx)
	if !ok {
		return "", false
	}
	return claims.Role, true
}

func GenerateJWT(userID, email, role string) (string, error) {

	expiry := config.GetConfig().GetInt("jwt.access_expiry")
	if expiry == 0 {
		expiry = 24
	}

	claims := &Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "kisaan-sathi",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID, email, role string) (string, error) {

	expiry := config.GetConfig().GetInt("jwt.refresh_expiry")
	if expiry == 0 {
		expiry = 24 * 30
	}

	claims := &Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "kisaan-sathi",
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateRefreshToken(token string) (*Claims, error) {

	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}