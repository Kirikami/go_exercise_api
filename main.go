package main

import (
	"github.com/kirikami/go_exercise_api/server"
	"github.com/kirikami/go_exercise_api/utils/application"
)

func main() {
	app := application.Application{}
	app.InitConfiguration()
	app.InitDatabase()
	server.StartServer(app)
}
