package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/ornellast/bucketeer/commons"
)

type Middleware = func(next http.Handler) http.Handler

func AcceptOnly(values ...string) Middleware {
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
					ctx := context.WithValue(r.Context(), commons.ContenTypeCtxKey, commons.CTypeJson)

					next.ServeHTTP(w, r.WithContext(ctx))
					return

				}

				if _, ok := acceptOnlyTypes[aType]; ok {

					ctx := context.WithValue(r.Context(), commons.ContenTypeCtxKey, aType)

					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			w.WriteHeader(http.StatusNotAcceptable)
		}
		return http.HandlerFunc(fn)
	}
}

func ContentTypeAutoSetter(defaultValue string) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cType := r.Context().Value(commons.ContenTypeCtxKey).(string)
			next.ServeHTTP(w, r)
			w.Header().Set(commons.CTypeHeader, cType)
		}
		return http.HandlerFunc(fn)
	}
}
