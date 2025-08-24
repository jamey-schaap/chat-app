package middleware

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	hj          http.Hijacker
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	rw := &responseWriterWrapper{ResponseWriter: w}
	if hj, ok := w.(http.Hijacker); ok {
		rw.hj = hj
	}
	return rw

}

func (rw *responseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if rw.hj == nil {
		return nil, nil, http.ErrNotSupported
	}
	return rw.hj.Hijack()
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
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
