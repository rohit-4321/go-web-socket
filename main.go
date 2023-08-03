package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const port = 8080

var upgrader = websocket.Upgrader{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello From GO server.")
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf("Upgrader err : ", err)
		return
	}
	defer c.Close()
	for {
		mt, p, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read : ", err)
			break

		}
		fmt.Printf("recv %s", p)
		err = c.WriteMessage(mt, []byte(string(p)+"From server.."))
		if err != nil {
			fmt.Println("Sending err:")
			break
		}
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/conn", socketHandler)
	fmt.Printf("Starting server in port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Error in staring server:\n", err)
	} else {
		fmt.Printf("Server serving on port %d\n", port)
	}

}
