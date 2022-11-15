package entities

import (
	"fmt"
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
			if dimVal != nil && reflect.TypeOf(dimVal).String() == "*proto.TensorShapeProto_Dimension_DimParam" {
				shape = append(shape, dim.GetDimParam())
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == "*proto.TensorShapeProto_Dimension_DimValue" {
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
			if dimVal != nil && reflect.TypeOf(dimVal).String() == "*proto.TensorShapeProto_Dimension_DimParam" {
				shape = append(shape, dim.GetDimParam())
			}
			if dimVal != nil && reflect.TypeOf(dimVal).String() == "*proto.TensorShapeProto_Dimension_DimValue" {
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

func (g *Graph) PrintGraphInput() {
	fmt.Println("======== Graph Inputs ========")
	fmt.Printf("Count inputs: %d\n", len(g.Inputs))

	for _, input := range g.Inputs {
		fmt.Printf("\n")
		fmt.Printf("Name: %s\n", input.Name)
		fmt.Printf("Type: %s\n", input.Type)
		fmt.Printf("Dimension: %d\n", input.Dimension)

		fmt.Printf("Shape: ")
		for i := 0; i < len(input.Shape)-1; i++ {
			fmt.Printf("%v * ", input.Shape[i])
		}
		fmt.Println(input.Shape[len(input.Shape)-1])
	}

	fmt.Println("==============================")
}

func (g *Graph) PrintGraphOutput() {
	fmt.Println("======== Graph Outputs ========")
	fmt.Printf("Count outputs: %d\n", len(g.Inputs))

	for _, output := range g.Outputs {
		fmt.Printf("\n")
		fmt.Printf("Name: %s\n", output.Name)
		fmt.Printf("Type: %s\n", output.Type)
		fmt.Printf("Dimension: %d\n", output.Dimension)

		fmt.Printf("Shape: ")
		for i := 0; i < len(output.Shape)-1; i++ {
			fmt.Printf("%v * ", output.Shape[i])
		}
		fmt.Println(output.Shape[len(output.Shape)-1])
	}

	fmt.Println("===============================")
}

func (g *Graph) PrintGraphNodes() {
	fmt.Println("======== Graph Nodes ========")
	fmt.Printf("Count nodes: %d\n", len(g.Nodes))

	for _, node := range g.Nodes {
		fmt.Printf("\n")
		fmt.Printf("Name: %s\n", node.Name)
		fmt.Printf("Type: %s\n", node.Type)

		fmt.Printf("Inputs: ")
		for _, input := range node.Input {
			fmt.Printf("%s ", input)
		}
		fmt.Printf("\n")

		fmt.Printf("Outputs: ")
		for _, output := range node.Output {
			fmt.Printf("%s ", output)
		}
		fmt.Printf("\n")
	}

	fmt.Println("=============================")
}

func (g *Graph) PrintGraph() {
	g.PrintGraphInput()
	g.PrintGraphOutput()
	g.PrintGraphNodes()
}
