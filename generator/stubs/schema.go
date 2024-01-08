package stubs

const (
	GetDataFunctionTemplate = `
func (m *{{ .ModelName }}) GetData(d *domain.{{ .ModelName }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ if not (or (eq .Name "ID") (eq .Name "Id")) }}{{ .Name | snakeToPascal }}: d.{{ .Name | snakeToPascal }},
		{{ end }}{{ end }}
	}
}
`
	SchemaStructTemplate = `
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
)
