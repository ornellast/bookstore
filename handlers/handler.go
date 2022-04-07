package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ornellast/bucketeer/commons"
	"github.com/ornellast/bucketeer/db"
	"github.com/ornellast/bucketeer/middlewares"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	middlewares.ConfigureMiddlewares(router)
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	itemsRoutes(router)
	return router
}

func setDefaultContentType(w http.ResponseWriter) {
	w.Header().Set(commons.CTypeHeader, commons.CTypeAppJson)
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
