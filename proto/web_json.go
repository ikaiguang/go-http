package protos

import (
	"fmt"
)

// WebJSON : web json
type WebJSON struct {
	Code    int32       `protobuf:"varint,1,opt,name=Code" json:"Code"`
	Message string      `protobuf:"bytes,2,opt,name=Message" json:"Message"`
	Data    interface{} `protobuf:"bytes,3,opt,name=Data" json:"Data"`
}

// Reset proto reset
func (w *WebJSON) Reset() {
	*w = WebJSON{}
}

// String proto string
func (w *WebJSON) String() string {
	// return proto.CompactTextString(m)
	return fmt.Sprintf("%#v", w)
}

// ProtoMessage proto message
func (w *WebJSON) ProtoMessage() {

}
