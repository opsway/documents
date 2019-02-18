package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
)

var Templates map[string]*Template

type TemplateData map[string]interface{}

type Template struct {
	path    string
	name    string
	index   *pongo2.Template
	Context map[string]interface{}
}

func (tmpl *Template) Render(data TemplateData, writer io.Writer) error {
	return tmpl.index.ExecuteWriter(pongo2.Context(data), writer)
}

func (tmpl *Template) loadTemplate() error {
	_, err := os.Stat(tmpl.path)

	if os.IsNotExist(err) {
		return err
	}

	tmpl.index, _ = pongo2.FromFile(tmpl.path)

	return nil
}

func NewTemplate(path string, name string) (*Template, error) {
	path = filepath.Join(path, name, "index.html")
	tmpl := &Template{path: path, name: name}

	return tmpl, tmpl.loadTemplate()
}

func GetTemplate(name string) (*Template, error) {
	tmpl, exists := Templates[name]
	if !exists {
		return nil, fmt.Errorf("template '%s' is not exist", name)

	}

	return tmpl, nil
}

func BuildTemplates(path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return fmt.Errorf("failed to read templates dir '%s'", path)
	}

	Templates = make(map[string]*Template)

	for _, file := range files {
		if file.IsDir() && file.Name()[0] != '.' {
			t, err := NewTemplate(path, file.Name())
			if err != nil {
				return err
			}
			Templates[file.Name()] = t
		}
	}

	return nil
}
