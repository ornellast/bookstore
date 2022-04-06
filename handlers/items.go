package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gitlab.com/ornellast/bucketeer/db"
	"gitlab.com/ornellast/bucketeer/models"
)

const itemIDKey = "itemId"

func items(router chi.Router) {
	router.Get("/", getAllItems)
	router.Post("/", createItem)
	router.Route("/{"+itemIDKey+"}", func(r chi.Router) {
		r.Use(ItemContext)
		r.Get("/", getItem)
		r.Put("/", updateItem)
		r.Delete("/", deleteItem)
	})
}

func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, itemIDKey)

		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Item ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)

		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
			// return
		}

		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		// render.Render(w, r, ErrBadRequest)
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)

	item, err := dbInstance.GetItemById(itemId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	itemData := models.Item{}

	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	item, err := dbInstance.UpdateItem(itemID, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)
	err := dbInstance.DeleteItem(itemId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
