package routes

import (
	"back/internal/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, authService *auth.Service) {
	mux.Handle("/auth", auth.OAuth2Handler(authService))
	mux.Handle("/auth/google/callback", auth.OAuth2CallbackHandler(authService))
}
