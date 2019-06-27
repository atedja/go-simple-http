package templates

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var htmls *sync.Map

func init() {
	htmls = &sync.Map{}
}

// Recursively reads all HTML files in the directory, and loads them as HTML templates
func Read(directory string) error {
	wd, _ := os.Getwd()
	directory = filepath.Join(wd, directory)

	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directory
		if info.IsDir() {
			return nil
		}

		fullPath := filepath.Join(directory, info.Name())
		data, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return err
		}

		name := info.Name()
		sdata := string(data)
		tmpl := template.New(name)
		tmpl, err = tmpl.Parse(sdata)
		if err != nil {
			return err
		}

		htmls.Store(name, tmpl)

		return nil
	})
}

// Loads an HTML template
func Load(name string) *template.Template {
	v, ok := htmls.Load(name)
	if !ok {
		return nil
	}

	var tmpl *template.Template
	switch v.(type) {
	case *template.Template:
		tmpl = v.(*template.Template)
	default:
		tmpl = nil
	}

	return tmpl
}
