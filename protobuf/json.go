package protobufs

import "fmt"

// WebJSON : web json
type WebJSON struct {
	Code    int32       `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

// Reset proto reset
func (w *WebJSON) Reset() {
	*w = WebJSON{}
}

// String proto string
func (w *WebJSON) String() string {
	return fmt.Sprintf("%#v", w)
}

// ProtoMessage proto message
func (w *WebJSON) ProtoMessage() {

}
