package main

import (
	"fmt"

	"github.com/nanozuki/akko"
	"github.com/nanozuki/akko/prop"
)

func main() {
	api := akko.OpenAPI("basic", "v1").
		Info(akko.Info().Description("basic server")).
		Path(
			akko.Path("/profile").
				GET(akko.Op("get_profile").
					Request(
						prop.Int("user_id"),
					).
					Response(
						prop.String("name"),
						prop.String("email"),
					)),
		)
	fmt.Println(api)
}
