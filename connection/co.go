package connection

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Co struct {
	Id        string
	Name      string
	Recipient *Co
	Conn      *websocket.Conn
}

var (
	connQueue  = make([]*Co, 0)
	queueMutex = sync.Mutex{}
)

func GetConn(w http.ResponseWriter, r *http.Request) (*Co, error) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Using bad Protocol")
		return nil, err
	}
	gCon := &Co{
		Recipient: nil,
		Name:      r.Header.Get("name"),
	}
	gCon.Id = uuid.New().String()
	gCon.Conn = c
	gCon.Conn.SetCloseHandler(handleCloseConnections)
	queueMutex.Lock()

	if len(connQueue) > 0 {
		recConn := connQueue[0]
		connQueue = connQueue[1:]
		gCon.Recipient = recConn
		recConn.Recipient = gCon

		gCon.Conn.WriteMessage(websocket.TextMessage, []byte("Connected with"+recConn.Name))
		recConn.Conn.WriteMessage(websocket.TextMessage, []byte("Connected with"+gCon.Name))

	} else {
		connQueue = []*Co{gCon}
	}
	queueMutex.Unlock()
	return gCon, nil
}

func (conn *Co) ReadLoop() {
	for {
		mt, p, err := conn.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Read Erro : ", err)
			break
		}
		haveReceipent := conn.Recipient
		if haveReceipent != nil {
			errSend := haveReceipent.Conn.WriteMessage(mt, p)
			if errSend != nil {
				fmt.Println("Error in sending to receiver...")
				break
			}
		}
	}
}
func (conn *Co) Close() {
	conn.Close()
}
func handleCloseConnections(code int, text string) error {
	fmt.Print("connection closed................")
	return nil
}