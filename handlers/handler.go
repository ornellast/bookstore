package handlers

import (
	"context"
	"net/http"
	"strings"

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
	router.Use(middleware.AllowContentType("application/json", "text/xml"))
	router.Use(acceptOnly("application/json", "application/xml"))

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	// router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(setContentType)
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

func acceptOnly(values ...string) Middleware {
	acceptOnlyTypes := make(map[string]struct{}, len(values))
	for _, ctype := range values {
		acceptOnlyTypes[strings.TrimSpace(strings.ToLower(ctype))] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			acceptValue := strings.ToLower(strings.TrimSpace(r.Header.Get("Accept")))
			if acceptValue == "" {
				// skip check for empty content body
				next.ServeHTTP(w, r)
				return
			}

			acceptArray := strings.Split(acceptValue, ",")

			for idx := 0; idx < len(acceptArray); idx++ {
				aType := acceptArray[idx]
				if i := strings.Index(aType, ";"); i > -1 {
					aType = aType[0:i]
				}

				if aType == "*/*" {
					ctx := context.WithValue(r.Context(), CTypeCtxKey, CTypeJson)

					next.ServeHTTP(w, r.WithContext(ctx))
					return

				}

				if _, ok := acceptOnlyTypes[aType]; ok {

					ctx := context.WithValue(r.Context(), CTypeCtxKey, aType)

					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			w.WriteHeader(http.StatusNotAcceptable)
		}
		return http.HandlerFunc(fn)
	}
}

func setContentType(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cType := r.Context().Value(CTypeCtxKey).(string)
		w.Header().Set(CTypeHeader, cType)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
