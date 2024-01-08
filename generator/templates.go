package generator

const (
	templateText = `
func (m *{{ .ModelName }}) Generate({{ range $index, $field := .Fields }}{{ if $index }}, {{ end }}{{ $field.Name | toCamelCase }} {{ if eq $field.Optional true }}*{{ end }} {{ $field.Type }}{{ end }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ .Name | snakeToPascal  }}: {{ .Name | toCamelCase }},
		{{ end }}
	}
}`
	getDataFunctionTemplate = `
func (m *{{ .ModelName }}) GetData(d *domain.{{ .ModelName }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ if not (or (eq .Name "ID") (eq .Name "Id")) }}{{ .Name | snakeToPascal }}: d.{{ .Name | snakeToPascal }},
		{{ end }}{{ end }}
	}
}
`
	domainStructTemplate = `
type {{ .StructName }} struct {
	{{ range .Fields }}{{ if .Optional }}{{ .Name | snakeToPascal }} *{{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\"`" + `{{ else }}{{ .Name | snakeToPascal }} {{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\"`" + `{{ end }}
	{{ end }}
}
`
	schemaStructTemplate = `
{{ if .GotID }}
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"{{ .ModuleName }}/domain"
)
{{ else }}
	import "{{ .ModuleName }}/domain"
{{ end }}

type {{ .StructName }} struct {
	{{ range .Fields }}{{ if .IsID }}{{ .Name | snakeToPascal }} primitive.ObjectID ` + "`json:\"id,omitempty\" bson:\"_id,omitempty\"`" + `{{ else }}{{ if .Optional }}{{ .Name | snakeToPascal }} *{{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\" bson:\"{{ .Name | toLower }},omitempty\"`" + `{{ else }}{{ .Name | snakeToPascal }} {{ .Type }} ` + "`json:\"{{ .Name | toLower }},omitempty\" bson:\"{{ .Name | toLower }},omitempty\"`" + `{{ end }}{{ end }}
{{ end }}
}
`
	interfaceHeaderTemplate = `
package repository

import (
	"context"
	{{ if .GotSson }}"{{.Module}}/pkg/sson"{{ end }}
	"{{.Module}}/domain"
)
`
)
