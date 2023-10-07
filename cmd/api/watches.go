package main

import (
	"fmt"
	"jewelry.abgdrv.com/internal/data"
	"net/http"
	"time"
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

	watch := data.Watch{
		ID:        id,
		CreatedAt: time.Now(),
		Brand:     "Rolex",
		Model:     "Submariner",
		DialColor: "Black",
		StrapType: "Stainless Steel",
		Price:     9999.99,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"watch": watch}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
