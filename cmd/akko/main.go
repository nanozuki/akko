package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/nanozuki/akko"
)

var (
	i   string
	o   string
	mod string
)

func init() {
	flag.StringVar(&i, "i", "", "input directory")
	flag.StringVar(&o, "o", "", "output directory")
	flag.StringVar(&mod, "mod", "", "module name")
}

func main() {
	flag.Parse()
	if err := os.MkdirAll(".akko", 0755); err != nil {
		log.Fatal(err)
	}

	cliFile, err := os.OpenFile(".akko/akko.go", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer cliFile.Close()

	if err := akko.Templates.ExecuteTemplate(cliFile, "command", akko.TmplCommandData{
		ServiceImportPath: filepath.Join(mod, i),
		ServicePkg:        filepath.Base(i),
	}); err != nil {
		log.Fatal(err)
	}
}
