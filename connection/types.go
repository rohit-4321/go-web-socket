package connection

import "encoding/json"

type OnConnectMessage struct {
	RecipientName string `json:"recipientName"`
	IsCaller      bool   `json:"isCaller"`
}

func (Cmeg *OnConnectMessage) GetJson() []byte {
	res, _ := json.Marshal(Cmeg)
	return res
}

func GetOnConnectMessage(recipientName string, isCaller bool) *OnConnectMessage {
	return &OnConnectMessage{
		RecipientName: recipientName,
		IsCaller:      isCaller,
	}
}
