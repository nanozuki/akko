package main

import (
	"github.com/nanozuki/ononoki"
	"github.com/nanozuki/ononoki/prop"
)

func main() {
	api := ononoki.OpenAPI("basic", "v1").
		Info(ononoki.Info().Description("basic server")).
		Path(
			ononoki.Path("/profile").GET(
				ononoki.Op("get_profile").
					Request(
						prop.Int("user_id"),
					).
					Response(
						prop.String("name"),
						prop.String("email"),
					),
			),
		)
	api.Run()
}
