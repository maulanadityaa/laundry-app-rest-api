package helper

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	appName                = os.Getenv("APP_NAME")
	jwtSecret              = os.Getenv("JWT_SECRET")
	jwtSignatureKey []byte = []byte(jwtSecret)
)

type Claims struct {
	AccountId string `json:"accountId"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(accountId, role, email string) (string, error) {
	claims := Claims{
		AccountId: accountId,
		Role:      role,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		tokenString = parts[1]
		fmt.Println(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSignatureKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			fmt.Println(err.Error())
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetJWTClaims(c *gin.Context) jwt.MapClaims {
	tokenString := c.GetHeader("Authorization")

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return nil
	}

	tokenString = parts[1]

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSignatureKey, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims
}
