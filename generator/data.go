package generator

// Data represents the parsed YAML data.
type Data struct {
	Module     string             `yaml:"module"`
	FileName   string             `yaml:"file_name"`
	Models     []*Model           `yaml:"models"`
	Interfaces []*InterfaceConfig `yaml:"interfaces"`
}
