package handlers

import "net/http"

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

const CTypeCtxKey = "CtxCType"

type Middleware = func(next http.Handler) http.Handler

func SetContentTypeHeader(w http.ResponseWriter, r *http.Request) {
	cType := r.Context().Value(CTypeCtxKey).(string)
	w.Header().Set(CTypeHeader, cType)
}
