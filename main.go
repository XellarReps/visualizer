package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"visualizer/pkg/entities"
	"visualizer/pkg/parser"
)

type commandArgs struct {
	inputPath  string
	writeMode  string
	infoMode   string
	outputPath string
}

type flagParams struct {
	name  string
	value string
	usage string
}

func main() {
	var args commandArgs

	params := map[*string]flagParams{
		&args.inputPath: {name: "input_path", value: "", usage: "input path (onnx file)"},
		&args.writeMode: {name: "write_mode", value: "stdout", usage: "output mode: stdout, file"},
		&args.infoMode: {name: "info_mode", value: "nodes", usage: "output filtering: graph (all information), " +
			"nodes (will return a list of nodes)"},
		&args.outputPath: {name: "output_path", value: "", usage: "output path"},
	}

	for ptr, param := range params {
		flag.StringVar(ptr, param.name, param.value, param.usage)
	}
	flag.Parse()

	if args.writeMode != "stdout" && args.writeMode != "file" {
		err := errors.New("error in command line argument 'write_mode', read --help")
		fmt.Println(err)
		return
	}

	if args.infoMode != "graph" && args.infoMode != "nodes" {
		err := errors.New("error in command line argument 'info_mode', read --help")
		fmt.Println(err)
		return
	}

	if args.inputPath == "" {
		err := errors.New("input path not specified")
		fmt.Println(err)
		return
	}

	model, err := parser.ReadModelFromFile(args.inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if args.writeMode == "file" && args.outputPath == "" {
		err := errors.New("output path not specified")
		fmt.Println(err)
		return
	}

	file := os.Stdout
	fileFlag := false
	if args.writeMode == "file" {
		file, err = os.OpenFile(args.outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileFlag = true
	}

	if args.infoMode == "graph" {
		modelInfo := entities.GetModelInfo(model)
		err := modelInfo.PrintModelInfo(file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	graphProto := model.GetGraph()
	graph := entities.GraphProtoToGraph(graphProto)
	err = graph.PrintGraph(file, args.infoMode)
	if err != nil {
		fmt.Println(err)
		return
	}
	if fileFlag {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
