package util

import (
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
)

const traceIDHeader = "x-trace-id"

// Tracer middleware function to generate a 'x-trace-id' header on all requests
func Tracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// checks the header existence. If it's not present, adds the header with a
		// random uuid
		traceID := r.Header.Get(traceIDHeader)
		if traceID == "" {
			traceID = uuid.New().String()
			r.Header.Set(traceIDHeader, traceID)
		}

		recorder := httptest.NewRecorder()
		next.ServeHTTP(recorder, r)

		w.Header().Set(traceIDHeader, traceID)
		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}

		w.Write(recorder.Body.Bytes())
	})
}
