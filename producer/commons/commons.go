package commons

const (
	CTypeHeader  = "Content-Type"
	AcceptHeader = "Accept"
)

const (
	CTypeAppJson = "application/json"
	CTypeText    = "text/plain"
	CTypeAppXml  = "application/xml"
	CTypeTxtXml  = "text/xml"
)

const Endpoint = "https://jsonplaceholder.typicode.com"

type ContextBaseKey struct {
	Name string
}

func (bk *ContextBaseKey) String() string {
	return "app/base/context/key/" + bk.Name
}

var AcceptContenTypeNegotiatedKey = &ContextBaseKey{Name: "Accept/Content-Type Accepted"}

type Identifier interface {
	Id() string
}
