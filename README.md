# Visualizer

## Description

This go module is designed to help get all the necessary information about the models of well-known frameworks 
(currently onnx)

At the moment, you can completely get the graph of the execution of the model as well as general information about the
model

### Compilation and launch

In the root of the project, run:

```bash
go build
./visualizer <filename>
```

### Command line arguments
The module supports different startup modes and support for command line arguments has been added for this.
To find out information about all the arguments, you need to type in the command line after compilation:
```bash
./visualizer --help
```

Detailed description of each argument:
```text
--input_path (string) the path to the input file of the .onnx format
--output_path (string) the path where the output file will lie if the argument write_mode=file
--write_mode (string) output mode, there are two modes: stdout (output to console) and file (output to file)
--info_mode (string) at the moment there are two modes of information representation: graph (all information about the graph) and nodes (enumeration of graph nodes)
```

### Example
The launch for example was carried out on resnet50. Onnx file taken from zenodo.\
Link to the file: [zenodo](https://zenodo.org/record/2592612/files/resnet50_v1.onnx)\
To run the example, it is necessary (in the root of the project):
```bash
mkdir inputs
mkdir output
cd inputs
wget https://zenodo.org/record/2592612/files/resnet50_v1.onnx
cd ..
go build
./visualizer --input_path=inputs/resnet50_v1.onnx --output_path=output/resnet_all.txt --write_mode=file --info_mode=graph
./visualizer --input_path=inputs/resnet50_v1.onnx --output_path=output/resnet_nodes.txt --write_mode=file --info_mode=nodes
```

## Links
The onnx.proto file was taken from the official [onnx repository](https://github.com/onnx/onnx)
