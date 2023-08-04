package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rohit-4321/go-web-socket/connection"
)

var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("name")
	if userName == "" {
		handleNoNameFound(w)
		return
	}
	c, err := connection.GetConn(w, r)
	if err != nil {
		return
	}
	defer c.Close()
	c.ReadLoop()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello From GO server.")
}

func handleNoNameFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprint(w, "No name found")
}
func handleCloseConnections(code int, text string) error {
	fmt.Print("connection closed................")
	return nil
}
