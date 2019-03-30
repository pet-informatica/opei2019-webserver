package api_gateway

import (
	"fmt"

	"github.com/labstack/echo"
)

const (
	EntryPoint = "/hello"

	CreateHelloRoute       = "/"
	CreateHelloRouteHAL    = "/"
	CreateHelloRouteFormat = "/"
)

type HelloAPI struct {
}

func MakeHelloRoutes(g *echo.Group) {
	api := HelloAPI{}

	g.GET(CreateHelloRoute, api.HelloHandler)
}

func (api HelloAPI) HelloHandler(ctx echo.Context) error {
	fmt.Println("Hello world!")
	return nil
}
