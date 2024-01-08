package stubs

const (
	GenerateFunc = `
func (m *{{ .ModelName }}) Generate({{ range $index, $field := .Fields }}{{ if $index }}, {{ end }}{{ $field.Name | toCamelCase }} {{ if eq $field.Optional true }}*{{ end }} {{ $field.Type }}{{ end }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ .Name | snakeToPascal  }}: {{ .Name | toCamelCase }},
		{{ end }}
	}
}`
	DomainStructTemplate = `
type {{ .StructName }} struct {
	{{ range .Fields }}{{ if .Optional }}{{ .Name | snakeToPascal }} *{{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\"`" + `{{ else }}{{ .Name | snakeToPascal }} {{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\"`" + `{{ end }}
	{{ end }}
}
`
)
