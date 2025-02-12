package main

import (
	"fmt"
	"net/http"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Log the site/origin that made the request
	logRequestOrigin(r)

	// Define a secure cookie with security flags
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "secure123", // Changed the value for clarity
		Path:     "/",
		HttpOnly: true,                    // Prevent access via JavaScript
		Secure:   true,                    // Allow transmission only over HTTPS
		SameSite: http.SameSiteStrictMode, // Restrict to same-site requests
	})

	// HTML response
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Set Cookie</title>
	</head>
	<body>
		<h1>Secure cookie set! Reload the page to send the cookie to the server.</h1>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, html)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	// Log the site/origin that made the request
	logRequestOrigin(r)

	// Read the cookie from the request
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// If the cookie is not present or invalid
		http.Error(w, "Forbidden: Missing or invalid session_token", http.StatusForbidden)
		return
	}

	// Validate the cookie value
	if cookie.Value != "secure123" {
		http.Error(w, "Forbidden: Invalid session_token", http.StatusForbidden)
		return
	}

	// Respond with a protected resource
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Protected Resource</title>
	</head>
	<body>
		<h1>Welcome to the protected resource!</h1>
		<p>Your session is valid.</p>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, html)
}

// Function to log the site/origin of the request
func logRequestOrigin(r *http.Request) {
	origin := r.Header.Get("Origin")
	referer := r.Header.Get("Referer")

	fmt.Println("Request Origin:")
	if origin != "" {
		fmt.Printf("- Origin: %s\n", origin)
	}
	if referer != "" {
		fmt.Printf("- Referer: %s\n", referer)
	}
	if origin == "" && referer == "" {
		fmt.Println("- Origin and Referer headers are missing.")
	}
}

func main() {
	// Route to set the cookie
	http.HandleFunc("/set", setCookieHandler)

	// Route for the protected resource
	http.HandleFunc("/protected-resource", resourceHandler)

	// Start the server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
