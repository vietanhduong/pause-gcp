package protoutil

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	legacyproto "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

var (
	unmarshaler       = jsonpb.Unmarshaler{AllowUnknownFields: true}
	strictUnmarshaler = jsonpb.Unmarshaler{}
)

func Unmarshal(b []byte, m proto.Message) error {
	return strictUnmarshaler.Unmarshal(bytes.NewReader(b), legacyproto.MessageV1(m))
}

func UnmarshalAllowUnknown(b []byte, m proto.Message) error {
	return unmarshaler.Unmarshal(bytes.NewReader(b), legacyproto.MessageV1(m))
}
