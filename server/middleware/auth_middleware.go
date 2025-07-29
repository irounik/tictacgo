package middleware

import (
	"context"
	"log"
	"net/http"
	"tictacgo/server/config"
	"tictacgo/server/service"
)

type AuthMiddleware interface {
	AuthrizeReqeust(next http.Handler) http.Handler
}

type SessionAuthMiddleware struct {
	sessionService service.SessionService
}

func NewSessionAuthMiddleware(sessionService service.SessionService) *SessionAuthMiddleware {
	return &SessionAuthMiddleware{sessionService: sessionService}
}

func (mw *SessionAuthMiddleware) AuthrizeReqeust(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(config.SESSION_TOKEN)
		if err != nil {
			http.Error(w, "User not logged in", http.StatusUnauthorized)
			return
		}

		sessionToken := cookie.Value
		session, err := mw.sessionService.GetByToken(sessionToken)

		if err != nil {
			http.Error(w, "Session expired, please login again!", http.StatusUnauthorized)
			return
		}

		log.Printf("Authenticated reqeuest on endpoint: %s", r.Pattern)
		ctx := context.WithValue(r.Context(), config.AUTH_USERNAME, session.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
