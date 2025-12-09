package main

import (
	"14-TestingAPI/configs"
	"14-TestingAPI/internal/auth"
	"14-TestingAPI/internal/link"
	"14-TestingAPI/internal/stat"
	"14-TestingAPI/internal/user"
	"14-TestingAPI/pkg/db"
	"14-TestingAPI/pkg/event"
	"14-TestingAPI/pkg/middleware"
	"fmt"
	"log"
	"net/http"
)

// func tickOperation(ctx context.Context) {
// 	ticker := time.NewTicker(200 * time.Millisecond)
// 	defer ticker.Stop() // Остановим ticker, когда функция завершится
// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("tick:", time.Now())
// 		case <-ctx.Done():
// 			fmt.Println("Cancel")
// 			return
// 		}
// 	}
// }

// func main2() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go tickOperation(ctx)

// 	time.Sleep(3 * time.Second)

//		cancel()
//		time.Sleep(1 * time.Second)
//	}
func main() {

	conf := configs.LoadConfig()

	// router := http.NewServeMux()
	// hello.NewHelloHandler(router)
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)
	//Services
	AuthService := auth.NewAuthService(userRepository)
	StatService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: AuthService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	go StatService.AddClick()

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
