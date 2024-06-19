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
	Email string
	jwt.RegisteredClaims
}

func JWTGenerate(username string) (string, error) {
	// expired one hour
	expTime := time.Now().Add(time.Second * 60 * 60)

	claims := JWTClaim{
		Email: username,
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

		// go to handler
		ctx.Next()
	}
}
