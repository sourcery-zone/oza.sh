package render

import (
	"bytes"
	"text/template"

	"github.com/sourcery-zone/oza.sh/internal/collectors"
)

// RenderStatus render the format string, given the context from the path.
func RenderStatus(format, path string) (string, error) {
	tmpl := template.New("status").Funcs(template.FuncMap{
		"Git": func() *collectors.Git {
			return collectors.NewGit(path)
		},
	})

	tmpl, err := tmpl.Parse(format)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return "", err
	}

	return buf.String(), nil
}
