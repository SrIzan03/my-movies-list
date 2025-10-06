package main

import (
	"back/internal/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/auth", http.HandlerFunc(auth.OAuth2Handler))
	mux.Handle("/auth/google/callback", http.HandlerFunc(auth.OAuth2CallbackHandler))
}
