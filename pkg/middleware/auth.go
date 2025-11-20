package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			log.Println("Bearer токена нет !!!")
			//next.ServeHTTP(w, r)
			//r.Response.StatusCode = 500
			return
		}
		authorization = strings.TrimPrefix(authorization, "Bearer ")
		log.Println(authorization)
		next.ServeHTTP(w, r)
	})
}
