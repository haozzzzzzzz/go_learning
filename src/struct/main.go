package main

import "fmt"

type Identify struct {
	id uint32
}

func (m *Identify) getID() uint32 {
	return m.id
}

type Message struct {
	strMsg  string
}

func (m *Message) getMessage() string {
	return m.strMsg
}

type Wrapper struct {
	Message
	Identify
}

func (m *Wrapper) getMessage() string {
	return m.strMsg + ", world"
}

func main() {
	wrapper := Wrapper{
		Message: Message{
			strMsg: "hello",
		},
		Identify: Identify{
			id: 123,
		},
	}
	fmt.Println(wrapper.Message.getMessage())
	fmt.Println(wrapper.getMessage())
	fmt.Println(wrapper.getID())
}
