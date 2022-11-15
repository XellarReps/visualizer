package entities

import (
	"fmt"
	"visualizer/pkg/proto"
)

type ModelInfo struct {
	IrVersion       int64
	ModelVersion    int64
	ProducerName    string
	ProducerVersion string
	Domain          string
}

func GetModelInfo(modelProto *proto.ModelProto) ModelInfo {
	return ModelInfo{
		IrVersion:       modelProto.GetIrVersion(),
		ModelVersion:    modelProto.GetModelVersion(),
		ProducerName:    modelProto.GetProducerName(),
		ProducerVersion: modelProto.GetProducerVersion(),
		Domain:          modelProto.GetDomain(),
	}
}

func (m *ModelInfo) PrintModelInfo() {
	fmt.Println("======== Model Info ========")
	fmt.Printf("IR version: %d\n", m.IrVersion)
	fmt.Printf("Model version: %d\n", m.ModelVersion)
	fmt.Printf("Producer name: %s\n", m.ProducerName)
	fmt.Printf("Producer version: %s\n", m.ProducerVersion)
	fmt.Printf("Domain: %s\n", m.Domain)
	fmt.Println("=============================")
}
