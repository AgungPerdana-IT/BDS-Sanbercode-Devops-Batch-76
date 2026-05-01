package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Golang Backend 🚀\nini app terbaru coba auto deploy\n🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server running on port %s\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
