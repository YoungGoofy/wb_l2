package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}
