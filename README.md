# Visualizer

## Description

This go module is designed to help get all the necessary information about the models of well-known frameworks 
(currently onnx)

At the moment, you can completely get the graph of the execution of the model as well as general information about the
model

To run the code, you need to:
- create a directory:
```bash
mkdir install
```
- put a file with the extension .onnx to the /inputs directory;
- compile and run the program with the command line argument, which is the name of the onnx model file.

### Compilation and launch

In the root of the project, run:

```bash
go build
./visualizer <filename>
```

## Links
The onnx.proto file was taken from the official [onnx repository](https://github.com/onnx/onnx)

Also thanks to [Viktor Scherbakov](https://github.com/ViktorooReps) for providing an example of the onnx model