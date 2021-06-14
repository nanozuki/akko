package main

import (
	"github.com/nanozuki/ononoki"
	"github.com/nanozuki/ononoki/desc"
	"github.com/nanozuki/ononoki/prop"
)

func main() {
	ononoki.Resquest(
		prop.String("name").In(desc.InPath),
	)
	ononoki.Response(
		prop.Object("detail",
			prop.String("name"),
		),
	)
}
