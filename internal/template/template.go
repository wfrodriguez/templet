package template

import (
	"bytes"
	"regexp"
	"text/template"
)

type VarMap map[string]string

var funcs = template.FuncMap{
	// Some functions...
}

// GetVariables devuelve un map con las variables que se van a utilizar en la plantilla
func GetVariableNames(tpl string) ([]string, error) {
	reVal, err := regexp.Compile(`(?i)\{\{(?:[^\.]+)?\.(\w+)(?:[^\}]+)?\}\}`)
	if err != nil {
		return nil, err
	}
	rt := make([]string, 0)
	res := reVal.FindAllStringSubmatch(tpl, -1)
	for _, s := range res {
		rt = append(rt, s[1])
	}

	return unique(rt), nil
}

// Render se encarga de crear el texto asociado a un archivo e incluir las variables definidas
func Render(tpl string, vars VarMap) (string, error) {
	tmpl := template.New("Template")
	tmpl, err := tmpl.Parse(tpl)
	if err != nil {
		return "", err
	}
	tmpl = tmpl.Funcs(funcs)

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, vars)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// unique es una función que devuelve un slice sin elementos repetidos
func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
