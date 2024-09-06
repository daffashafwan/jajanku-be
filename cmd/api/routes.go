package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	mux := mux.NewRouter()

	mux.NotFoundHandler = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)
	mux.Use(app.authenticate)

	mux.HandleFunc("/status", app.status).Methods("GET")

	mux.HandleFunc("/users", app.createUser).Methods("POST")
	mux.HandleFunc("/authentication-tokens", app.createAuthenticationToken).Methods("POST")

	authenticatedRoutes := mux.NewRoute().Subrouter()
	authenticatedRoutes.Use(app.requireAuthenticatedUser)
	authenticatedRoutes.HandleFunc("/protected", app.protected).Methods("GET")

	protectedRoutes := mux.NewRoute().Subrouter()
	protectedRoutes.Use(app.requireBasicAuthentication)
	protectedRoutes.HandleFunc("/basic-auth-protected", app.protected).Methods("GET")

	return mux
}
