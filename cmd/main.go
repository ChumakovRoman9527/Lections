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
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop() // Остановим ticker, когда функция завершится
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick:", time.Now())
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go tickOperation(ctx)

	time.Sleep(3 * time.Second)

	cancel()
	time.Sleep(1 * time.Second)
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
