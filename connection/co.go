package connection

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	connQueue  = make([]*Co, 0)
	queueMutex = sync.Mutex{}
	upgrader   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Add this line if you want to allow any origin to connect
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type Co struct {
	Id        string
	Name      string
	Recipient *Co
	Conn      *websocket.Conn
	IsCaller  bool
}

func (conn *Co) SendMessage(messageType int, data []byte) error {
	err := conn.Conn.WriteMessage(messageType, data)
	if err != nil {
		return err
	}
	return nil
}

func (conn *Co) ReadLoop() {
	for {
		mt, p, err := conn.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Read Erro : ", err)
			break
		}
		if conn.Recipient != nil {
			errSend := conn.Recipient.SendMessage(mt, p)
			if errSend != nil {
				fmt.Println("Error in sending to receiver...")
				break
			}
		}
	}
}
func (co *Co) Close() {
	co.Conn.Close()
}
func GetConn(w http.ResponseWriter, r *http.Request) (*Co, error) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		fmt.Fprint(w, "Using bad Protocol")
		return nil, err
	}
	query := r.URL.Query()
	userName := query.Get("name")
	gCon := &Co{
		Recipient: nil,
		Name:      userName,
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

		gCon.IsCaller = false
		recConn.IsCaller = true
		gCon.SendMessage(websocket.TextMessage, GetOnConnectMessage(recConn.Name, gCon.IsCaller).GetJson())
		recConn.SendMessage(websocket.TextMessage, GetOnConnectMessage(gCon.Name, recConn.IsCaller).GetJson())

	} else {
		connQueue = []*Co{gCon}
	}
	queueMutex.Unlock()

	return gCon, nil
}

func handleCloseConnections(code int, text string) error {
	fmt.Print("connection closed................")
	return nil
}
