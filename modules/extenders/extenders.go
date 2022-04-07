package extenders

import (
	"net/http"
)

type Extenders struct {
	name string
}

func NewExtenders(name string) *Extenders {
	return &Extenders{name: name}
}

func (e *Extenders) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (e *Extenders) Lookup(r *http.Request) (view string, data interface{}, has bool) {
	return e.name, nil, true
}
