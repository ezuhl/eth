package routes

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

//setup expected interfaces for router
type Handler interface {
	SetAuthToken(*jwtauth.JWTAuth)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetApiKey(w http.ResponseWriter, r *http.Request)
	GetAverageHeight(w http.ResponseWriter, r *http.Request)
}
