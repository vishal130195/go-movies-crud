package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the current directory (ui folder)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Create a file server handler
	fs := http.FileServer(http.Dir(dir))

	// Handle all requests
	http.Handle("/", fs)

	port := "3000"
	fmt.Printf("🌐 Starting UI server...\n")
	fmt.Printf("📍 Serving files from: %s\n", dir)
	fmt.Printf("🚀 UI Server running at: http://localhost:%s\n", port)
	fmt.Printf("📱 Open in browser: http://localhost:%s/index.html\n", port)
	fmt.Printf("🛑 Press Ctrl+C to stop\n\n")

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
