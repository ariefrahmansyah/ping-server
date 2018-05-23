package main

import (
	"fmt"
	"net/http"

	"github.com/ariefrahmansyah/ping"
)

func main() {
	fmt.Println("🏓 http://localhost:8080/ping")

	http.HandleFunc("/ping", ping.Handler)
	http.ListenAndServe(":8080", nil)
}
