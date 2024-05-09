package main

import (
	"fmt"
	"net/http"
	"strings"
)

func protecaoInjecao(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			sanitizedUsername := sanitizeInput(username)
			sanitizedPassword := sanitizeInput(password)

			fmt.Println("Username sanitizado:", sanitizedUsername)
			fmt.Println("Password sanitizado:", sanitizedPassword)
		}

		next.ServeHTTP(w, r)
	})
}

func sanitizeInput(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "'", "")
	input = strings.ReplaceAll(input, "\"", "")
	return input
}
