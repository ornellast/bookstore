package commons

import "fmt"

type BookTopic int

const (
	BookCreate BookTopic = iota
	BookUpdate
	BookDelete
	BookPatch
)

func (bt BookTopic) String() string {
	if bt < 0 || bt > 3 {
		vl := fmt.Sprintf("%d", bt)
		panic(vl + " is't a valid value")
	}
	return []string{"bookstore-new-book", "bookstore-update-book", "bookstore-delete-book", "bookstore-patch-book"}[bt]

}

// const (
