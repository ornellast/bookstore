package models

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	XMLName     xml.Name `json:"-" xml:"book"`
	ID          int      `json:"id" xml:"id,attr"`
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description"`
	CreatedAt   string   `json:"created_at" xml:"createdAt"`
}
type BookList struct {
	XMLName xml.Name `json:"-" xml:"books"`
	Books   []Book   `json:"books" xml:"book"`
}

func (i *Book) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*BookList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (bk Book) Id() string {
	return strconv.Itoa(bk.ID)
}
