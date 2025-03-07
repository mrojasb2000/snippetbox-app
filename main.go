package main

import (
	"log"
	"net/http"

	"snippetboxapp.org/controllers"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// Register the two new handler functions and corresponding URL patterns
	// with the servemux, in exactly the same way we did before.
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/snippet", controllers.ShowSnippet)
	http.HandleFunc("/snippet/create", controllers.CreateSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// in two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)

}
