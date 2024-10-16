package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// NoCheck does nothing except calling the asynchronous handler
// Use it as an example
func NoCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, start.Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, headslice := range r.Header["Authorization"] {
			if headslice == "allowed" {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusForbidden) // Utiliser le code 403 pour une interdiction
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"message": "forbidden"}); err != nil {
			w.WriteHeader(http.StatusInternalServerError) // En cas d'erreur avec l'encodage JSON
		}
	})
}
