package middleware

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (r *responseWriter) WriteHeader(code int) {
	if r.wroteHeader {
		return
	}

	r.status = code
	r.ResponseWriter.WriteHeader(code)
	r.wroteHeader = true
	return
}

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errorMsg := fmt.Sprintf("%v", err)
					logger.Error(errorMsg)
				}
			}()

			start := time.Now()
			wrappedWriter := wrapResponseWriter(w)
			next.ServeHTTP(wrappedWriter, r)

			requestDuration := time.Since(start)
			logMsg := fmt.Sprintf("%d %s %s %v", wrappedWriter.status, r.Method, r.URL.Path, requestDuration)
			logger.Info(logMsg, zap.Int("Status", wrappedWriter.status), zap.String("Method", r.Method), zap.String("Path", r.URL.Path), zap.Duration("Duration", requestDuration))
		})
	}
}
