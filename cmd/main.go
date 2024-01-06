package main

import (
	"template/file"
	"template/generator"
	"template/yaml"
)

func main() {
	generate()
}

func generate() {
	fileNames, err := file.ExtractFileNamesFromDirectory("configs")
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {
		data, err := yaml.ParseYaml(fileName)
		if err != nil {
			panic(err)
		}
		generator.Run(data)
	}
}
