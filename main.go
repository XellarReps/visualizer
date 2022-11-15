package main

import (
	"fmt"
	"os"
	"visualizer/pkg/entities"
	"visualizer/pkg/parser"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("file name not specified (provide a command line argument)")
		return
	}

	model, err := parser.ReadModelFromFile("inputs/" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	modelInfo := entities.GetModelInfo(model)
	modelInfo.PrintModelInfo()

	graphProto := model.GetGraph()
	graph := entities.GraphProtoToGraph(graphProto)
	graph.PrintGraph()
}
