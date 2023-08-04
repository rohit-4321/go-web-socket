package main

import (
	"fmt"
	"net/http"

	"github.com/rohit-4321/go-web-socket/handler"
)

const port = 8080

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/conn", handler.SocketHandler)
	fmt.Printf("Starting server in port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		fmt.Println("Error in staring server:\n", err)
	} else {
		fmt.Printf("Server serving on port %d\n", port)
	}

}
