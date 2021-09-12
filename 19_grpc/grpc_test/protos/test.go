package protos

//go:generate protoc -I ../../grpc_test/protos -I $GOPATH/include --go_out=../../ test.proto

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// CreateHello ...
func CreateHello(name string) *Hello {
	return &Hello{
		Name: name,
		Time: ptypes.TimestampNow(),
	}
}

// UnmarshalHello ...
func UnmarshalHello(data []byte) (*Hello, error) {
	ret := &Hello{}

	if err := proto.Unmarshal(data, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// MarshalHello ...
func MarshalHello(data *Hello) ([]byte, error) {
	return proto.Marshal(data)
}
