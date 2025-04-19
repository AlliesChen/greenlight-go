package main

import (
	"fmt"
	"net/http"
)

// POST /v1/movies
func (app *application) craeteMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "show the movie details of movie %d\n", id)
}
