package app

import "encoding/json"

type Message[T any] struct {
	body T
}

func NewMessage[T any](body T) *Message[T] {
	return &Message[T]{body: body}
}

func (m *Message[T]) Body() T {
	return m.body
}

func (m *Message[T]) RawBody() ([]byte, error) {
	return json.Marshal(m.body)
}
