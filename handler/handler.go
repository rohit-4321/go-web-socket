package handler

import (
	"fmt"
	"net/http"

	"github.com/rohit-4321/go-web-socket/connection"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("name")
	println("Request")
	if userName == "" {
		handleNoNameFound(w)
		return
	}
	c, err := connection.GetConn(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	c.ReadLoop()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello From Go server.")
}

func handleNoNameFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprint(w, "No name found")
}
