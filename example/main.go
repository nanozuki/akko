package main

import (
	"github.com/nanozuki/ononoki"
	"github.com/nanozuki/ononoki/prop"
	"github.com/nanozuki/ononoki/typ"
)

func main() {
	ononoki.Resquest(
		prop.String("name").InPath(),
	)
	ononoki.Response(
		prop.Object("detail", typ.Object("UserDetail", prop.String("name"))),
	)
}
