package yaml

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"template/generator"
)

func ParseYaml(filePath string) (*generator.Data, error) {
	fullPath := filepath.Join("configs", filePath)

	yamlContent, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	var data *generator.Data
	err = yaml.Unmarshal(yamlContent, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
