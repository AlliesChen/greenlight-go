package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlliesChen/greenlight-go/internal/data"
)

// POST /v1/movies
func (app *application) craeteMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Inception",
		Runtime:   102,
		Genres:    []string{"Action", "Sci-Fi"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
