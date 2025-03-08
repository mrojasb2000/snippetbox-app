package controllers

import "net/http"

// Define a home function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
