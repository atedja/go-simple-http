package templates

import (
	"html/template"
	"os"
	"path/filepath"
)

var All *template.Template

func init() {
	wd, _ := os.Getwd()
	directory := filepath.Join(wd, "templates", "*")
	All = template.Must(template.ParseGlob(directory))
}
