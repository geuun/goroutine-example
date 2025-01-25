package utils

import (
	"net/http"
	"strings"
)

type DynamicHandler struct {
	routes map[string]http.HandlerFunc
}

func NewDynamicHandler() *DynamicHandler {
	return &DynamicHandler{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (h *DynamicHandler) AddRoute(path string, handler http.HandlerFunc) {
	h.routes[path] = handler
}

func (h *DynamicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	for pattern, handler := range h.routes {
		if matchRoute(pattern, path) {
			handler(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func matchRoute(pattern, path string) bool {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := 0; i < len(patternParts); i++ {
		if !strings.HasPrefix(patternParts[i], ":") && patternParts[i] != pathParts[i] {
			return false
		}
	}

	return true
}
