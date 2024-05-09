package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	maxFailedLoginAttempts = 3
	failedLoginAttempts    = make(map[string]int)
	failedLoginAttemptsMux sync.Mutex
)

func deteccaoIntrusao(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" && r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			if username == "admin" && password == "password123" {
				next.ServeHTTP(w, r)
				return
			} else {
				failedLoginAttemptsMux.Lock()
				failedLoginAttempts[r.RemoteAddr]++
				count := failedLoginAttempts[r.RemoteAddr]
				failedLoginAttemptsMux.Unlock()

				if count > maxFailedLoginAttempts {
					fmt.Printf("Intrusão detectada: Tentativa de login falha do IP %s\n", r.RemoteAddr)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
