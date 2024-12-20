package main

import (
	"fmt"
	"net/http"
	"time"

	"green_api.com/internal/data"
	"green_api.com/internal/validator"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "ABC",
		RunTime:   102,
		Genres:    []string{"drama", "romance"},
		Version:   1,
	}

	err = app.wrtieJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		RunTime int32    `json:"runtime"`
		Year    int32    `json:"year"`
		Generes []string `json:"genres"`
	}
	v := validator.New()

	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(input.RunTime != 0, "runtime", "must be provided")
	v.Check(input.RunTime > 0, "runtime", "must be a positive integer")
	v.Check(input.Generes != nil, "genres", "must be provided")
	v.Check(len(input.Generes) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Generes) <= 5, "genres", "must not contain more than 5 genres")

	err := app.readJson(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v", input)
}
