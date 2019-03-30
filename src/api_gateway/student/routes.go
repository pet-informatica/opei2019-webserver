package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opei2019-webservice/src/api_gateway/util"
	std "opei2019-webservice/src/student"
	"opei2019-webservice/src/student_service"

	"github.com/bugsnag/bugsnag-go"
	"github.com/labstack/echo"
)

// Create all string routes here
const (
	EntryPoint = "/students"

	CreateStudentRoute       = ""
	CreateStudentRouteHAL    = ""
	CreateStudentRouteFormat = ""

	GetStudentByIdRoute       = ""
	GetStudentByIdRouteHAL    = "?id={id}"
	GetStudentByIdRouteFormat = "?id=%s"

	contentType = "application/hal+json"
)

type StudentAPI struct {
	StudentService student_service.Service
}

// MakeStudentRoutes creates all the routes for students
// POST / - Create a new student
// GET - Get the student data
func MakeStudentRoutes(g *echo.Group, studentService student_service.Service) {
	api := StudentAPI{studentService}

	g.POST(CreateStudentRoute, api.CreateStudentHandler)
	g.GET(GetStudentByIdRoute, api.GetStudentHandler)
}

func (api *StudentAPI) CreateStudentHandler(ctx echo.Context) error {
	c := util.GetContext(ctx)
	req := ctx.Request()

	student := &std.Student{}
	if err := json.NewDecoder(req.Body).Decode(student); err != nil {
		bugsnag.Notify(err) // Send an event error to bugsnag
		return &echo.HTTPError{Code: http.StatusUnprocessableEntity, Message: "Failed to read student json"}
	}

	status, student, err := api.StudentService.CreateStudent(c, student)
	switch {
	case err != nil:
		bugsnag.Notify(err)
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Student service failed"}
	case status != http.StatusCreated:
		bugsnag.Notify(err)
		return &echo.HTTPError{Code: int(status), Message: "Failed to create student"}
	}

	body, _ := json.Marshal(student)
	return ctx.Blob(int(status), contentType, body)
}

func (api *StudentAPI) GetStudentHandler(ctx echo.Context) error {
	id := ctx.QueryParam("id")

	fmt.Println("É desta forma que pegamos uma query string. ID: ", id)
	return nil
}
