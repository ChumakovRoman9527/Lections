package main

import (
	"11-JWTAUTH/configs"
	"11-JWTAUTH/internal/auth"
	"11-JWTAUTH/internal/link"
	"11-JWTAUTH/internal/user"
	"11-JWTAUTH/pkg/db"
	"11-JWTAUTH/pkg/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {

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
