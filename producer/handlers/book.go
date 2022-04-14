package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ornellast/bookstore/producer/commons"
	"github.com/ornellast/bookstore/producer/db"
	"github.com/ornellast/bookstore/producer/kafkaprd"
	"github.com/ornellast/bookstore/producer/models"
)

const bookIDParam = "bookId"

var ctxBookIDKey = fmt.Sprint(&commons.ContextBaseKey{Name: "book/id"})

func booksRoutes(r *chi.Mux) {
	r.Route("/books", func(router chi.Router) {
		router.Get("/", getAllBooks)
		router.Post("/", createBook)
		router.Route("/{"+bookIDParam+"}", func(r chi.Router) {
			r.Use(BookContext)
			r.Get("/", getBook)
			r.Put("/", createChangesNotifier(commons.BookUpdate))
			r.Patch("/", createChangesNotifier(commons.BookPatch))
			r.Delete("/", deleteBook)
		})
	})
}

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookId := chi.URLParam(r, bookIDParam)

		if bookId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Book ID is required")))
			return
		}
		id, err := strconv.Atoi(bookId)

		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid book ID")))
			// return
		}

		ctx := context.WithValue(r.Context(), ctxBookIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dbInstance.GetAllBooks()

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, books); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := render.Bind(r, book); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := kafkaprd.NewEvent(commons.BookCreate.String(), book); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.Context().Value(ctxBookIDKey).(int)

	book, err := dbInstance.GetBookById(bookId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func createChangesNotifier(topic commons.BookTopic) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookData := models.Book{}

		if err := render.Bind(r, &bookData); err != nil {
			render.Render(w, r, ErrBadRequest)
			return
		}

		bookData.ID = r.Context().Value(ctxBookIDKey).(int)

		if err := kafkaprd.NewEvent(topic.String(), bookData); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}

		if err := render.Render(w, r, &bookData); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.Context().Value(ctxBookIDKey).(int)

	if err := kafkaprd.NewEvent(commons.BookDelete.String(), &models.Book{
		ID: bookId,
	}); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	w.WriteHeader(204)

	/*
		err := dbInstance.DeleteBook(bookId)
		if err != nil {
			if err == db.ErrNoMatch {
				render.Render(w, r, ErrNotFound)
			} else {
				render.Render(w, r, ServerErrorRenderer(err))
			}
			return
		} */
}
