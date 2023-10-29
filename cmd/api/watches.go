package main

import (
	"fmt"
	"jewelry.abgdrv.com/internal/data"
	"jewelry.abgdrv.com/internal/validator"
	"net/http"
	"time"
)

// POST "/v1/watches"
func (app *application) createWatchHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Brand     string  `json:"brand"`
		Model     string  `json:"model,omitempty"`
		DialColor string  `json:"dial_color"`
		StrapType string  `json:"strap_type"`
		Diameter  int64   `json:"diameter"`
		Energy    string  `json:"energy"`
		Gender    string  `json:"gender"`
		Price     float64 `json:"price"`
		ImageURL  string  `json:"image_url"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	watch := &data.Watch{
		Brand:     input.Brand,
		Model:     input.Model,
		DialColor: input.DialColor,
		StrapType: input.StrapType,
		Diameter:  input.Diameter,
		Energy:    input.Energy,
		Gender:    input.Gender,
		Price:     input.Price,
		ImageURL:  input.ImageURL,
	}

	v := validator.New()

	if data.ValidateWatch(v, watch); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Watches.Insert(watch)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/watches/%d", watch.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"watch": watch}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input)
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
		Diameter:  40,
		Energy:    "Mechanical",
		Gender:    "Male",
		Price:     9999.99,
		ImageURL:  "https://shorturl.at/qDKN1",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"watch": watch}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
