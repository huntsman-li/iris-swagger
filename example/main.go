package main

import (
	"github.com/huntsman-li/iris-swagger"
	"github.com/huntsman-li/iris-swagger/swaggerFiles"
	"github.com/kataras/iris"

	_ "github.com/huntsman-li/iris-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := iris.New()

	r.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	r.Run(iris.Addr(":8080"))
}
