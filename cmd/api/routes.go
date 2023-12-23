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

	router.HandlerFunc(http.MethodGet, "/v1/watches",
		app.requirePermission("watches:read", app.listWatchesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/watches",
		app.requirePermission("watches:write", app.createWatchHandler))
	router.HandlerFunc(http.MethodGet, "/v1/watches/:id",
		app.requirePermission("watches:read", app.showWatchHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/watches/:id",
		app.requirePermission("watches:write", app.updateWatchHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/watches/:id",
		app.requirePermission("watches:write", app.deleteWatchHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
