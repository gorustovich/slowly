package app

import (
	"net/http"
	"mime"
	"context"
	"time"
	"sync"
	"bytes"
)

const timeout = time.Second * 5

func ValidateJsonType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "" {
			NewErrorResponse("content-Type header must be application/json", http.StatusBadRequest).Write(w)
			return
		}
		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			NewErrorResponse("malformed Content-Type header", http.StatusBadRequest).Write(w)
			return
		}

		if mt != "application/json" {
			NewErrorResponse("content-Type header must be application/json", http.StatusUnsupportedMediaType).Write(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CancelTooLong(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		r = r.WithContext(ctx)
		done := make(chan struct{})
		tw := &timeoutWriter{
			w: w,
			h: make(http.Header),
		}
		go func() {
			next.ServeHTTP(tw, r)
			close(done)
		}()
		select {
		case <-done:
			tw.mu.Lock()
			defer tw.mu.Unlock()
			dst := w.Header()
			for k, vv := range tw.h {
				dst[k] = vv
			}
			if !tw.wroteHeader {
				tw.code = http.StatusOK
			}
			w.WriteHeader(tw.code)
			w.Write(tw.wbuf.Bytes())
		case <-ctx.Done():
			tw.mu.Lock()
			defer tw.mu.Unlock()
			NewErrorResponse("timeout too long", http.StatusBadRequest).Write(w)
		}
	})
}

type timeoutWriter struct {
	w    http.ResponseWriter
	h    http.Header
	wbuf bytes.Buffer

	mu          sync.Mutex
	wroteHeader bool
	code        int
}

func (tw *timeoutWriter) Header() http.Header { return tw.h }

func (tw *timeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if !tw.wroteHeader {
		tw.writeHeader(http.StatusOK)
	}
	return tw.wbuf.Write(p)
}

func (tw *timeoutWriter) WriteHeader(code int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}
