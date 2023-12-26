package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yumu-uw/simple-ipam-server/api"

	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	simpleIpam := api.NewSimpleIpam()

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, simpleIpam)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal((e.Start(net.JoinHostPort("0.0.0.0", *port))))
}
