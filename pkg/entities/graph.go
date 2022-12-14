package entities

import (
	"fmt"
	"os"
	"reflect"
	"visualizer/pkg/proto"
)

type InputOutputGraph struct {
	Name      string
	Type      string
	Dimension int
	Shape     []any
}

type Node struct {
	Name   string
	Type   string
	Input  []string
	Output []string
}

type Graph struct {
	Inputs  []InputOutputGraph
	Outputs []InputOutputGraph
	Nodes   []Node
}

const (
	dimValueType = "*proto.TensorShapeProto_Dimension_DimValue"
	dimParamType = "*proto.TensorShapeProto_Dimension_DimParam"
)

func GraphProtoToGraph(graphProto *proto.GraphProto) Graph {
	var graph Graph

	var inputs []InputOutputGraph
	for _, input := range graphProto.Input {
		inputType := input.GetType()
		tensorType := inputType.GetTensorType()
		elemType := proto.TensorProto_DataType_name[tensorType.GetElemType()]
		dims := tensorType.GetShape().GetDim()

		var shape []any
		for _, dim := range dims {
			dimVal := dim.GetValue()
			if dimVal == nil {
				shape = append(shape, "?")
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == dimParamType {
				shape = append(shape, dim.GetDimParam())
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == dimValueType {
				shape = append(shape, dim.GetDimValue())
			}
		}

		inputs = append(inputs, InputOutputGraph{
			Name:      input.GetName(),
			Type:      elemType,
			Dimension: len(dims),
			Shape:     shape,
		})
	}
	graph.Inputs = inputs

	var outputs []InputOutputGraph
	for _, output := range graphProto.Output {
		outputType := output.GetType()
		tensorType := outputType.GetTensorType()
		elemType := proto.TensorProto_DataType_name[tensorType.GetElemType()]
		dims := tensorType.GetShape().GetDim()

		var shape []any
		for _, dim := range dims {
			dimVal := dim.GetValue()
			if dimVal == nil {
				shape = append(shape, "?")
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == dimParamType {
				shape = append(shape, dim.GetDimParam())
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == dimValueType {
				shape = append(shape, dim.GetDimValue())
			}
		}

		outputs = append(outputs, InputOutputGraph{
			Name:      output.GetName(),
			Type:      elemType,
			Dimension: len(dims),
			Shape:     shape,
		})
	}
	graph.Outputs = outputs

	var nodes []Node
	for _, node := range graphProto.Node {
		var inputArray, outputArray []string
		for _, input := range node.Input {
			inputArray = append(inputArray, input)
		}
		for _, output := range node.Output {
			outputArray = append(outputArray, output)
		}

		nodes = append(nodes, Node{
			Name:   node.GetName(),
			Type:   node.GetOpType(),
			Input:  inputArray,
			Output: outputArray,
		})
	}
	graph.Nodes = nodes

	return graph
}

func (g *Graph) PrintGraphInput(file *os.File) error {
	if _, err := file.WriteString("======== Graph Inputs ========\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Count inputs: %d\n", len(g.Inputs))); err != nil {
		return err
	}

	for _, input := range g.Inputs {
		if _, err := file.WriteString("\n"); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Name: %s\n", input.Name)); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Type: %s\n", input.Type)); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Dimension: %d\n", input.Dimension)); err != nil {
			return err
		}

		if _, err := file.WriteString("Shape: "); err != nil {
			return err
		}
		for i := 0; i < len(input.Shape)-1; i++ {
			if _, err := file.WriteString(fmt.Sprintf("%v * ", input.Shape[i])); err != nil {
				return err
			}
		}
		if _, err := file.WriteString(fmt.Sprintf("%v\n", input.Shape[len(input.Shape)-1])); err != nil {
			return err
		}
	}

	_, err := file.WriteString("==============================\n")
	return err
}

func (g *Graph) PrintGraphOutput(file *os.File) error {
	if _, err := file.WriteString("======== Graph Outputs ========\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(fmt.Sprintf("Count outputs: %d\n", len(g.Inputs))); err != nil {
		return err
	}

	for _, output := range g.Outputs {
		if _, err := file.WriteString("\n"); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Name: %s\n", output.Name)); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Type: %s\n", output.Type)); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Dimension: %d\n", output.Dimension)); err != nil {
			return err
		}

		if _, err := file.WriteString("Shape: "); err != nil {
			return err
		}
		for i := 0; i < len(output.Shape)-1; i++ {
			if _, err := file.WriteString(fmt.Sprintf("%v * ", output.Shape[i])); err != nil {
				return err
			}
		}
		if _, err := file.WriteString(fmt.Sprintf("%v\n", output.Shape[len(output.Shape)-1])); err != nil {
			return err
		}
	}

	_, err := file.WriteString("===============================\n")
	return err
}

func (g *Graph) PrintGraphNodes(file *os.File, mode string) error {
	if mode == "graph" {
		if _, err := file.WriteString("======== Graph Nodes ========\n"); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Count nodes: %d\n", len(g.Nodes))); err != nil {
			return err
		}
	}

	for _, node := range g.Nodes {
		if mode == "graph" {
			if _, err := file.WriteString("\n"); err != nil {
				return err
			}
		}
		if _, err := file.WriteString(fmt.Sprintf("Name: %s\n", node.Name)); err != nil {
			return err
		}
		if _, err := file.WriteString(fmt.Sprintf("Type: %s\n", node.Type)); err != nil {
			return err
		}

		if mode == "graph" {
			if _, err := file.WriteString("Inputs: "); err != nil {
				return err
			}
			for _, input := range node.Input {
				if _, err := file.WriteString(fmt.Sprintf("%s ", input)); err != nil {
					return err
				}
			}
			if _, err := file.WriteString("\n"); err != nil {
				return err
			}

			if _, err := file.WriteString("Outputs: "); err != nil {
				return err
			}
			for _, output := range node.Output {
				if _, err := file.WriteString(fmt.Sprintf("%s ", output)); err != nil {
					return err
				}
			}
			if _, err := file.WriteString("\n"); err != nil {
				return err
			}
		}
	}

	_, err := file.WriteString("=============================\n")
	return err
}

func (g *Graph) PrintGraph(file *os.File, mode string) error {
	if mode == "graph" {
		err := g.PrintGraphInput(file)
		if err != nil {
			return err
		}
	}
	if mode == "graph" {
		err := g.PrintGraphOutput(file)
		if err != nil {
			return err
		}
	}
	err := g.PrintGraphNodes(file, mode)
	return err
}
