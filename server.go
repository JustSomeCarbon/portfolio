package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
}
*/

func main() {

	// define a new file server rooted in the static directory
	var fs http.Handler = http.FileServer(http.Dir("./static"))

	// pass the file server as the handler function for the server
	// all static files must be prepended with the path "/static/"
	// we must strip the "/static/" prefix from the file server path
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	start_server(":8080")
}

// dictates what port the server is listening on and starts the server
func start_server(port string) {
	fmt.Printf("Server listening at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
