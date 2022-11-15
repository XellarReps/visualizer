package parser

import (
	protobuf "google.golang.org/protobuf/proto"
	"os"
	"visualizer/pkg/proto"
)

func ReadModelFromFile(filename string) (*proto.ModelProto, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	modelProto := new(proto.ModelProto)
	err = protobuf.Unmarshal(bytes, modelProto)
	if err != nil {
		return nil, err
	}

	return modelProto, nil
}
