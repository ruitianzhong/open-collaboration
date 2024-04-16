package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/auth/login" && r.URL.Path != "/api/user/check-if-login" {
				w.Header().Del("X_USER_ID")

				session, _ := store.Get(r, "dm-session")
				if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
					log.Println("reject")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				userId, ok := session.Values["username"].(string)
				if !ok {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.Header().Set("X_USER_ID", userId)
				next.ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}

			next.ServeHTTP(w, r)

		})
}
