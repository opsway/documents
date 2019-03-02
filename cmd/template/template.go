package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
)

var templates map[string]*Template

// Context refers various context of templates
type Context map[string]interface{}

// Template is structure of template
type Template struct {
	path    string
	name    string
	index   *pongo2.Template
	Context map[string]interface{}
}

// Render outputs the templates to writer
func (tmpl *Template) Render(context Context, writer io.Writer) error {
	return tmpl.index.ExecuteWriter(pongo2.Context(context), writer)
}

func (tmpl *Template) load() error {
	_, err := os.Stat(tmpl.path)

	if os.IsNotExist(err) {
		return err
	}

	tmpl.index, err = pongo2.FromFile(tmpl.path)

	if err != nil {
		return err
	}

	return nil
}

// NewTemplate make template specified dir path and name of template
func NewTemplate(path string, name string) (*Template, error) {
	path = filepath.Join(path, name, "index.html")
	tmpl := &Template{path: path, name: name}

	err := tmpl.load()
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

// GetTemplate returns template specified by name
func GetTemplate(name string) (*Template, error) {
	tmpl, exists := templates[name]
	if !exists {
		return nil, fmt.Errorf("template '%s' is not exist", name)

	}

	return tmpl, nil
}

// BuildTemplates makes templates which is specified by dir path
func BuildTemplates(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return fmt.Errorf("failed to read templates dir '%s'", path)
	}

	templates = make(map[string]*Template)

	for _, file := range files {
		if file.IsDir() && file.Name()[0] != '.' {
			tmpl, err := NewTemplate(path, file.Name())
			if err != nil {
				return err
			}
			templates[file.Name()] = tmpl
		}
	}

	return nil
}
