package middlewares

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const secretEnv = "LAYNAS_SECRET"

var (
	ErrAuthorizationHeaderMissing = errors.New("authorization header missing")
	ErrAuthorizationHeaderFormat  = errors.New("authorization header has wrong format")
	ErrWrongSecret                = errors.New("wrong secret")
)

func WithAuthentication() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		if authorization == "" {
			_ = ctx.AbortWithError(http.StatusBadRequest, ErrAuthorizationHeaderMissing)
		}

		secret, err := base64.StdEncoding.DecodeString(authorization)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, ErrAuthorizationHeaderFormat)
		}

		if string(secret) != os.Getenv(secretEnv) {
			_ = ctx.AbortWithError(http.StatusUnauthorized, ErrWrongSecret)
		}
	}
}
