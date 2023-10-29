package main

import (
	"errors"
	"fmt"
	"jewelry.abgdrv.com/internal/data"
	"jewelry.abgdrv.com/internal/validator"
	"net/http"
	"strconv"
)

// POST "/v1/watches"
func (app *application) createWatchHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Brand     string  `json:"brand"`
		Model     string  `json:"model,omitempty"`
		DialColor string  `json:"dial_color"`
		StrapType string  `json:"strap_type"`
		Diameter  int8    `json:"diameter"`
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

	watch, err := app.models.Watches.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"watch": watch}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateWatchHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	watch, err := app.models.Watches.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(watch.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	var input struct {
		Brand     *string  `json:"brand"`
		Model     *string  `json:"model,omitempty"`
		DialColor *string  `json:"dial_color"`
		StrapType *string  `json:"strap_type"`
		Diameter  *int8    `json:"diameter"`
		Energy    *string  `json:"energy"`
		Gender    *string  `json:"gender"`
		Price     *float64 `json:"price"`
		ImageURL  *string  `json:"image_url"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Brand != nil {
		watch.Brand = *input.Brand
	}
	if input.Model != nil {
		watch.Model = *input.Model
	}
	if input.DialColor != nil {
		watch.DialColor = *input.DialColor
	}
	if input.StrapType != nil {
		watch.StrapType = *input.StrapType
	}
	if input.Diameter != nil {
		watch.Diameter = *input.Diameter
	}
	if input.Energy != nil {
		watch.Energy = *input.Energy
	}
	if input.Gender != nil {
		watch.Gender = *input.Gender
	}
	if input.Price != nil {
		watch.Price = *input.Price
	}
	if input.ImageURL != nil {
		watch.ImageURL = *input.ImageURL
	}

	v := validator.New()
	if data.ValidateWatch(v, watch); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Watches.Update(watch)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"watch": watch}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Watches.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "watch successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listWatchesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Brand     string
		DialColor string
		Filters   data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Brand = app.readString(qs, "brand", "")
	input.DialColor = app.readString(qs, "dial_color", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{"id", "brand", "dial_color", "-id", "-brand", "-dial_color"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	watches, err := app.models.Watches.GetAll(input.Brand, input.DialColor, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"watches": watches}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input)
}
