package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// POST "/v1/watches"
func (app *application) createWatchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new watch")
}

// GET "/v1/watches/:id"
func (app *application) showWatchHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL params
	params := httprouter.ParamsFromContext(r.Context())

	// Get id param
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of watch %d\n", id)
}
