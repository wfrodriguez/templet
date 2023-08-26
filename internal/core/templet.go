package core

import (
	"os"
	"path/filepath"

	"github.com/wfrodriguez/templet/internal/template"
)

type Application struct {
	templateDir string
}

func NewApplication() (*Application, TempletError) {
	d, e := os.UserHomeDir()
	if e != nil {
		return nil, NewError(e, "HomeDir")
	}
	return &Application{
		templateDir: filepath.Join(d, ".config", "templet"),
	}, nil
}

func (a *Application) SetTemplateDir(dir string) TempletError {
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		return NewError(err, "Application.SetTemplateDir")
	}
	a.templateDir = dir

	return nil
}

func (a *Application) ExtractVariables(ext, name string) ([]string, TempletError) {
	body, err := a.GetFile(ext, name)
	if err != nil {
		return nil, err
	}

	vars, e := template.GetVariableNames(body)
	if e != nil {
		return nil, NewError(e, "Application.ExtractVariables")
	}

	return vars, nil
}

func (a *Application) GetFile(ext, name string) (string, TempletError) {
	fileName := filepath.Join(a.templateDir, ext, name+".tmpl")
	body, err := os.ReadFile(fileName)
	if err != nil {
		return "", NewError(err, "Application.GetFile")
	}

	return string(body), nil
}
