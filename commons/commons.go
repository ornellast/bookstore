package commons

const (
	CTypeHeader  = "Content-Type"
	AcceptHeader = "Accept"
)
const (
	CTypeJson = "application/json"
	CTypeText = "text/plain"
	CTypeXml  = "application/xml"
)

const Endpoint = "https://jsonplaceholder.typicode.com"

type contextBaseKey struct {
	Name string
}

func (bk *contextBaseKey) String() string {
	return "bucketeer/base/context/key/" + bk.Name
}

var ContenTypeCtxKey = &contextBaseKey{"CtxCType"}

// func SetContentTypeHeader(w http.ResponseWriter, r *http.Request) {
// 	cType := r.Context().Value(CTypeCtxKey).(string)
// 	w.Header().Set(CTypeHeader, cType)
// }
