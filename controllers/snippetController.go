package controllers

import "net/http"

// Add a showSnippet handler function.
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code
	// and the w.Write() method to write a "Method Not Allowed" response body.
	// We then from the function so that the subsequent code is not executed.
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add an 'Allow: POST' header to the response
		// header map. The first parameter is the header name, and the second parameter
		// is the header value.
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and "Method Not Allowed"
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
