package utils

import "encoding/json"

func GetByte(obj interface{}) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	return data
}

func GetString(obj interface{}) string {
	data, _ := json.Marshal(obj)

	return string(data)
}
