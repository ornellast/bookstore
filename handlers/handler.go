package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gitlab.com/ornellast/bucketeer/db"
)

const (
	cTypeHeader = "Content-Type"
	cTypeJson   = "application/json"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Route("/items", items)
	return router
}

func setDefaultContentType(w http.ResponseWriter) {
	w.Header().Set(cTypeHeader, cTypeJson)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultContentType(w)
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	setDefaultContentType(w)
	w.WriteHeader(404)
	render.Render(w, r, ErrNotFound)
}
