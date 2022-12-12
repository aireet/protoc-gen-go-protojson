// Code generated by protoc-core-go-proto-json. DO NOT EDIT.

package example

import (
	"fmt"
	"reflect"

	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func init() {

	// Use RegisterExampleTimestampCodec to register customize codec
	encoding.RegisterCodec(ExampleTimestampProtoJsonCodec{})

}

// Default codec
type ExampleTimestampProtoJsonCodec struct {
	protojson.MarshalOptions
	protojson.UnmarshalOptions
}

func (c ExampleTimestampProtoJsonCodec) Name() string {
	return "example.ExampleTimestamp"
}

// Customize codec
func RegisterExampleTimestampCodec(m protojson.MarshalOptions, u protojson.UnmarshalOptions) {
	encoding.RegisterCodec(ExampleTimestampProtoJsonCodec{m, u})
}

// MarshalJSON implements json.Marshaler
func (msg *ExampleTimestamp) MarshalJSON() ([]byte, error) {
	return encoding.GetCodec("example.ExampleTimestamp").Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ExampleTimestamp) UnmarshalJSON(b []byte) error {
	return encoding.GetCodec("example.ExampleTimestamp").Unmarshal(b, msg)
}

func (c ExampleTimestampProtoJsonCodec) Marshal(v interface{}) ([]byte, error) {
	if vv, ok := v.([]byte); ok {
		return vv, nil
	}
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
	}
	return c.MarshalOptions.Marshal(vv)
}

func (c ExampleTimestampProtoJsonCodec) Unmarshal(data []byte, v interface{}) error {
	if _, ok := v.(*[]byte); ok {
		reflect.ValueOf(v).Elem().SetBytes(data)
		return nil
	}
	vv, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
	}
	return c.UnmarshalOptions.Unmarshal(data, vv)
}
