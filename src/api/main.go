package main

import (
	"portfolio/infra/handler"
)

func main() {
	e := handler.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
