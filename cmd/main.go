package main

import (
	"12-Context/configs"
	"12-Context/internal/auth"
	"12-Context/internal/link"
	"12-Context/internal/user"
	"12-Context/pkg/db"
	"12-Context/pkg/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	type key string
	const EmailKey key = "Email"

	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, EmailKey, "A@A.ru") - нет ошибки
	// ctxWithValue := context.WithValue(ctx, "Not", "A@A.ru") //ошибка - такого ключа нету

	if userEmail, ok := ctxWithValue.Value(EmailKey).(string); ok {
		fmt.Println(userEmail)
	} else {
		fmt.Println("NoValue")
	}
}
func main2() {

	conf := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	db := db.NewDb(conf)
	router := http.NewServeMux()
	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	//Services
	AuthService := auth.NewAuthService(userRepository)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: AuthService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	//http.ListenAndServe(":8081", nil)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ошибка запуска сервера")
	}

}
