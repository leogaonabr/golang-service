package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

// Metrificator middleware func para gerar logs estruturados das requisições servidas pela API
func Metrificator(next http.Handler) http.Handler {
	hostname, _ := os.Hostname()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// copies the request body
		requestBytes, _ := ioutil.ReadAll(r.Body)
		bodyCopy := ioutil.NopCloser(bytes.NewBuffer(requestBytes))
		// puts the copy back at the request
		r.Body = bodyCopy

		// generates the log entry
		httpMetric := map[string]interface{}{
			"path":                 r.RequestURI,
			"protocol":             r.Proto,
			"method":               r.Method,
			"request_headers":      r.Header,
			"request_body":         string(requestBytes),
			"request_query_params": r.URL.Query(),
			"client_ip":            r.RemoteAddr,
		}

		logger := GetLogger().WithFields(map[string]interface{}{
			"metric": map[string]interface{}{
				"trace_id": r.Header.Get(traceIDHeader),
				"hostname": hostname,
				"module":   "api",
				"http":     httpMetric,
			},
		})
		defer logger.Info()

		start := time.Now()
		recorder := httptest.NewRecorder()
		next.ServeHTTP(recorder, r)
		latency := time.Since(start)

		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}

		responseBytes, _ := ioutil.ReadAll(recorder.Body)

		httpMetric["latency"] = (latency / time.Millisecond)
		httpMetric["status_code"] = recorder.Code
		httpMetric["response_body"] = string(responseBytes)

		w.Write(responseBytes)
	})
}
