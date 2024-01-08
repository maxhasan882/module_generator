package stubs

const (
	InterfaceHeaderTemplate = `
package repository

import (
	"context"
	{{ if .GotSson }}"{{.Module}}/pkg/sson"{{ end }}
	"{{.Module}}/domain"
)
`
)
