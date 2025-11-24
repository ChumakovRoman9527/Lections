package middleware

import (
	"12-Context/configs"
	"12-Context/pkg/jwt"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ConstEmailKey key = "ConstEmailKey"
)

func IsAuthed(next http.Handler, config *configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			log.Println("Bearer токена нет !!!")
			//next.ServeHTTP(w, r)
			//r.Response.StatusCode = 500
			return
		}
		token := strings.TrimPrefix(authorization, "Bearer ")
		isValid, data := jwt.NewJWT(config.Secret).Parse(token)
		fmt.Println(isValid)
		ctx := context.WithValue(r.Context(), ConstEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
