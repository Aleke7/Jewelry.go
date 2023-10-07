package main

import (
	"fmt"
	"net/http"
)

// POST "/v1/watches"
func (app *application) createWatchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new watch")
}

// GET "/v1/watches/:id"
func (app *application) showWatchHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of watch %d\n", id)
}
