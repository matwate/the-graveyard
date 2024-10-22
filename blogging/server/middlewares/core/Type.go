package core

import "net/http"

type OperationType int

const (
	Include OperationType = iota
	Exclude
)

type Middleware struct {
	Operation OperationType
	Func      func(next http.Handler) http.Handler
	Urls      map[string]struct{}
}

func (m *Middleware) Apply(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, ok := m.Urls[r.URL.Path]

		if (m.Operation == Include && ok) || (m.Operation == Exclude && !ok) {
			// Call the middleware function
			h = m.Func(h)
		}

		// Call the next handler in the chain
		h.ServeHTTP(w, r)
	})
}

func CreateStack_(xs ...Middleware) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			h = xs[i].Apply(h)
		}
		return h
	}
}
