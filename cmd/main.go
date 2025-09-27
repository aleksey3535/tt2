package main

import "tt2/internal/app"

// @title Blueprint Swagger API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name Aleksey
// @contact.email tuderinovav@mail.ru



// @BasePath /

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		panic(err)
	}
}