package router

import (
	"net/http"
	"tictacgo/server/handler"
	"tictacgo/server/middleware"
)

func NewRouter(authHandler handler.AuthHandler, authMiddleware middleware.AuthMiddleware, gameHandler handler.GameHandler) http.Handler {
	mux := http.NewServeMux()

	registerRoute := func(pattern string, handlers map[string]func(http.ResponseWriter, *http.Request)) {
		mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			if handler, exists := handlers[r.Method]; exists {
				handler(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}

	registerRoute("/signup", POST(authHandler.SignUp))
	registerRoute("/signin", POST(authHandler.SignIn))
	registerRoute("/logout", POST(authHandler.LogOut))

	mux.Handle("/game", authMiddleware.AuthrizeReqeust(http.HandlerFunc(gameHandler.NewGame)))
	return mux
}

func POST(postMethodHandler func(w http.ResponseWriter, r *http.Request)) map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		http.MethodPost: postMethodHandler,
	}
}
