package web

import (
	"fmt"
	"net/http"
)

type MapHandler struct {
	handlers map[string]func(c *Context)
}

func (h *MapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		ctx := NewContext(w, r)
		handler(ctx)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not Found"))
	}
}

func (h *MapHandler) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
