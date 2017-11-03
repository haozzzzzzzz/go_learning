package main

import (
	"encoding/json"
	"fmt"
)

type ResponseData map[string]interface{}
type ResponseMessage struct {
	Code uint32       `json:"code"`
	Msg  uint32       `json:"msg"`
	Data ResponseData `json:"data"`
}

func NewResponseMessage() *ResponseMessage {
	return &ResponseMessage{
		Data: make(ResponseData),
	}
}

func main() {
	msg := NewResponseMessage()
	byteJson, err := json.Marshal(msg)
	fmt.Println(string(byteJson), err)
}
