package handler

import (
	"net/http"
	"os"
)

var (
	adminLogin    = os.Getenv("ADMIN_LOGIN")
	adminPassword = os.Getenv("ADMIN_PASSWORD")
)

func (h *Handler) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != adminLogin || pass != adminPassword {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
