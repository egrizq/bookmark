package helpers

import (
	"auth-go/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

func JWTGenerate(username string) (string, error) {
	// expired 15 minute
	expTime := time.Now().Add(15 * time.Minute)

	claims := JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-token",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare tokens
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	log.Println("Success generate jwt")

	return token, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil || err == http.ErrNoCookie {
			model.Response(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		claims := &JWTClaim{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		log.Println("username from jwt:", claims.Username)

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				model.Response(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			model.Response(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if !token.Valid {
			model.Response(ctx, http.StatusInternalServerError, "token is invalid")
			return
		}
		ctx.Set("username", claims.Username)

		// go to handler
		ctx.Next()
	}
}
