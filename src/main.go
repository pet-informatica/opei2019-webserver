package main

import (
	"opei2019-webservice/src/api_gateway"
	"opei2019-webservice/src/api_gateway/student"
	"opei2019-webservice/src/student_service"

	"github.com/bugsnag/bugsnag-go"
	"github.com/labstack/echo"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          "7980aca4551b2d5420ce8048e46c5066",
		ProjectPackages: []string{"main", "opei2019-webservice"},
	})

	service := student_service.NewService()

	app := echo.New()

	helloRoute := app.Group(api_gateway.EntryPoint)
	api_gateway.MakeHelloRoutes(helloRoute)

	studentRoutes := app.Group(student.EntryPoint)
	student.MakeStudentRoutes(studentRoutes, service)

	app.Start(":8081")
}
