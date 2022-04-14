package models

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Item struct {
	XMLName     xml.Name `json:"-" xml:"item"`
	ID          int      `json:"id" xml:"id,attr"`
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description"`
	CreatedAt   string   `json:"created_at" xml:"createdAt"`
}
type ItemList struct {
	XMLName xml.Name `json:"-" xml:"items"`
	Items   []Item   `json:"items" xml:"item"`
}

func (i *Item) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
