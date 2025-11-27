package middleware

import (
	"13-AdvancedDB/configs"
	"13-AdvancedDB/pkg/jwt"
	"context"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ConstEmailKey key = "ConstEmailKey"
)

func writeUnAuthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
func IsAuthed(next http.Handler, config *configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			log.Println("Bearer токена нет !!!")
			writeUnAuthed(w)
			return
		}
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeUnAuthed(w)
			return
		}
		token := strings.TrimPrefix(authorization, "Bearer ")
		isValid, data := jwt.NewJWT(config.Secret).Parse(token)
		if !isValid {
			writeUnAuthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ConstEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
