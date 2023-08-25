package core

import "os"

type Application struct {
	TemplateDir string
}

func NewApplication() *Application {
	d, e := os.UserHomeDir()
	if e != nil {
		//
	}
	return &Application{
		TemplateDir: d,
	}
}
