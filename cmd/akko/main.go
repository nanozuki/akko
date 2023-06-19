package main

import (
	"flag"
	"fmt"

	"github.com/nanozuki/akko"
)

var (
	i   string
	o   string
	s   string
	mod string
)

func init() {
	flag.StringVar(&i, "i", "", "input directory")
	flag.StringVar(&o, "o", "", "output directory")
	flag.StringVar(&s, "s", "", "struct name")
	flag.StringVar(&mod, "mod", "", "module name")
}

// reflect
// func main() {
// 	flag.Parse()
// 	if err := os.MkdirAll(".akko", 0755); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	cliFile, err := os.OpenFile(".akko/akko.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cliFile.Close()
//
// 	if err := akko.Templates.ExecuteTemplate(cliFile, "command", akko.TmplCommandData{
// 		ServiceImportPath: filepath.Join(mod, i),
// 		ServicePkg:        filepath.Base(i),
// 		PkgDir:            i,
// 	}); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if err := exec.Command("go", "fmt", ".akko/akko.go").Run(); err != nil {
// 		log.Fatal(fmt.Errorf("failed to format .akko/akko.go: %w", err))
// 	}
// 	out, err := exec.Command("go", "run", ".akko/akko.go").Output()
// 	if err != nil {
// 		log.Fatal(fmt.Errorf("failed to run .akko/akko.go: %w", err))
// 	}
// 	fmt.Println("output:", string(out))
// }

// parse
func main() {
	flag.Parse()
	si, err := akko.ParseService(i, s)
	fmt.Printf("parse result: %+v, err: %v\n", si, err)
}
