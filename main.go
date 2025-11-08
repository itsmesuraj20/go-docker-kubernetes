package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define HTTP handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/message", messageHandler)

	// Start server
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go + Docker + Kubernetes</title>
		<style>
			body { font-family: Arial; max-width: 800px; margin: 50px auto; padding: 20px; }
			h1 { color: #00ADD8; }
			.card { background: #f4f4f4; padding: 20px; border-radius: 5px; margin: 20px 0; }
		</style>
	</head>
	<body>
		<h1>🚀 Welcome to Go + Docker + Kubernetes!</h1>
		<div class="card">
			<h2>System Info</h2>
			<p><strong>Hostname:</strong> %s</p>
			<p><strong>Port:</strong> %s</p>
		</div>
		<div class="card">
			<h2>Available Endpoints</h2>
			<ul>
				<li><a href="/">/</a> - Home page (this page)</li>
				<li><a href="/health">/health</a> - Health check endpoint</li>
				<li><a href="/api/message">/api/message</a> - API endpoint</li>
			</ul>
		</div>
	</body>
	</html>
	`
	hostname, _ := os.Hostname()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Fprintf(w, html, hostname, port)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy", "service": "go-k8s-demo"}`)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, `{"message": "Hello from Kubernetes!", "hostname": "%s"}`, hostname)
}