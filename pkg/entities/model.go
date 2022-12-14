package entities

import (
	"fmt"
	"os"
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

func (m *ModelInfo) PrintModelInfo(file *os.File) error {
	if _, err := file.WriteString("======== Model Info ========\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("IR version: %d\n", m.IrVersion)); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Model version: %d\n", m.ModelVersion)); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Producer name: %s\n", m.ProducerName)); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Producer version: %s\n", m.ProducerVersion)); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Domain: %s\n", m.Domain)); err != nil {
		return err
	}
	if _, err := file.WriteString("======== Model Info ========\n"); err != nil {
		return err
	}
	return nil
}
