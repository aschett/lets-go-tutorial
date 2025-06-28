package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aschett/lets-go-tutorial/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{
		Snippets: snippets,
	})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.tmpl", templateData{
		Snippet: snippet,
	})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form to create new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O go"
	content := "O go \n what a beatiful language \n u are"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
