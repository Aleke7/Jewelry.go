package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	// Init new httprouter
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/watches", app.listWatchesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/watches", app.createWatchHandler)
	router.HandlerFunc(http.MethodGet, "/v1/watches/:id", app.showWatchHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/watches/:id", app.updateWatchHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/watches/:id", app.deleteMovieHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	return app.recoverPanic(app.rateLimit(router))
}
