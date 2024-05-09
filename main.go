package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("struct"))

	middleware := protecaoDDoS(deteccaoIntrusao(protecaoInjecao(fs)))

	http.Handle("/", middleware)

	fmt.Println("Servidor ativo em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
