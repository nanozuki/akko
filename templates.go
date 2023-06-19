package akko

import (
	"embed"
	"text/template"
)

//go:embed templates/*.gotmpl
var tmplFs embed.FS

var Templates = template.Must(template.ParseFS(tmplFs, "templates/*.gotmpl"))

type TmplCommandData struct {
	ServiceImportPath string
	ServicePkg        string
	PkgDir            string
}
