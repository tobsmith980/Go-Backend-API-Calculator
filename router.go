package backendapicalculator

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "3000"

	router := http.NewServeMux()

	// TODO: Add handler functions
	router.HandleFunc("POST /add")
	router.HandleFunc("POST /subtract")
	router.HandleFunc("POST /multiply")
	router.HandleFunc("POST /divide")

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Sprintf("API Listening on port: %d", port)

	server.ListenAndServe()
}
