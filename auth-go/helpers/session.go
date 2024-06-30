package helpers

import (
	"auth-go/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var STORE *sessions.CookieStore

func InitSession() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("SESSIONS_KEY")
	STORE = sessions.NewCookieStore([]byte(key))
}

func CheckSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := STORE.Get(ctx.Request, "session")
		if err != nil {
			model.Response(ctx, http.StatusInternalServerError, err.Error())
			ctx.Abort()
			return
		}

		username, ok := session.Values["username"].(string)
		if !ok || username == "" {
			model.Response(ctx, http.StatusUnauthorized, "unauthorized user")
			ctx.Abort()
			return
		}
		log.Println("access for", username)

		expired := 3600
		STORE.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   expired, // 1 hour
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		}

		err = STORE.Save(ctx.Request, ctx.Writer, session)
		if err != nil {
			log.Println("err save session:", err.Error())
			model.Response(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		ctx.Next()
	}
}
