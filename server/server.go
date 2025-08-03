package server

import (
	"fmt"
	"log"
	"net/http"
	"tictacgo/server/handler"
	"tictacgo/server/middleware"
	"tictacgo/server/repository"
	"tictacgo/server/router"
	"tictacgo/server/service"
)

func Start(port int) {
	appRouter := setUpAppRouter()
	fmt.Printf("Starting server on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), appRouter)
	if err != nil {
		log.Fatalf("Failed to start server on port %d; %v", port, err)
	}
}

func setUpAppRouter() http.Handler {
	// Auth Flow
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	sessionService := service.NewSessionService()
	authHandler := handler.NewAuthHandler(*authService, sessionService)

	// Middleware
	authMiddleware := middleware.NewSessionAuthMiddleware(sessionService)

	// Game Flow
	gameRepo := repository.NewGameRepository()
	gameService := service.NewGameService(gameRepo)
	gameHandler := handler.NewGameHandler(gameService)

	// App router
	appRouter := router.NewRouter(authHandler, authMiddleware, gameHandler)
	return appRouter
}
