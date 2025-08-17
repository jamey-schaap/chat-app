package server

import "net/http"

func CorsMiddleware(handler http.Handler) http.Handler {
	return handler
}
