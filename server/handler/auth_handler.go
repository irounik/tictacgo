package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tictacgo/server/config"
	"tictacgo/server/model/request"
	"tictacgo/server/service"
	"time"
)

type AuthHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
}

type SessionAuthHandler struct {
	authService    service.AuthService
	sessionService service.SessionService
}

func NewAuthHandler(authService service.AuthService, sessionService service.SessionService) *SessionAuthHandler {
	return &SessionAuthHandler{authService: authService, sessionService: sessionService}
}

func (h *SessionAuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var authRequest request.AuthRequest
	decoder.Decode(&authRequest)

	// Validate request
	if err := validateAuthRequest(authRequest); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	err := h.authService.SignUp(authRequest)
	if err != nil {
		// TODO: Send proper err response codes
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	cookie, err := h.getSessionCookie(authRequest.Username)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong!"))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(200)
	w.Write([]byte("Sign up successful!"))
}

func (h *SessionAuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var authRequest request.AuthRequest
	decoder.Decode(&authRequest)

	// Validate request
	if err := validateAuthRequest(authRequest); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	err := h.authService.SignIn(authRequest)
	if err != nil {
		// TODO: Send proper err response codes
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	cookie, err := h.getSessionCookie(authRequest.Username)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong!"))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(200)
	w.Write([]byte("Signed in successfully!"))
}

func (h *SessionAuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(config.SESSION_TOKEN)
	if err != nil {
		http.Error(w, "User not logged in!", http.StatusUnauthorized)
	}

	sessionToken := cookie.Value
	h.sessionService.Invalidate(sessionToken)
	w.WriteHeader(200)
	w.Write([]byte("Logged out successfully!"))
}

func (h *SessionAuthHandler) getSessionCookie(username string) (*http.Cookie, error) {
	session, err := h.sessionService.Create(username)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     config.SESSION_TOKEN,
		Value:    session.Token,
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expires in 24 hours
		HttpOnly: true,                           // Prevents client-side script access
		Path:     "/",                            // Available for all paths
	}

	return cookie, nil
}

func validateAuthRequest(request request.AuthRequest) error {
	if request.Username == "" {
		return fmt.Errorf("username is required")
	}

	if len(request.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if request.Password == "" {
		return fmt.Errorf("password is required")
	}

	if len(request.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}

	return nil
}
