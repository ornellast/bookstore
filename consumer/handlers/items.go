package handlers

import (
	"net/http"

	"github.com/ornellast/bookstore/consumer/models"
)

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()

}

func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}

}

func getItem(w http.ResponseWriter, r *http.Request) {

	item, err := dbInstance.GetItemById(itemId)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	itemData := models.Item{}
	item, err := dbInstance.UpdateItem(itemID, itemData)
}

func patchItem(w http.ResponseWriter, r *http.Request) {
	itemData := models.Item{}
	itemDB, err := dbInstance.GetItemById(itemID)

	if itemData.Description == "" {
		itemData.Description = itemDB.Description
	}

	item, err := dbInstance.UpdateItem(itemID, itemData)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	err := dbInstance.DeleteItem(itemId)
}
