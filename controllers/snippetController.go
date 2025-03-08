package controllers

import "net/http"

// Add a showSnippet handler function.
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
