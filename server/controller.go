package server

import "github.com/labstack/echo/v4"

type RequestMethod int

const (
	GET RequestMethod = iota
	POST
	PUT
	DELETE
)

type EchoController interface {
	GetMethod() RequestMethod
	GetPath() string
	GetHandler() func(c echo.Context) error
}

type Controller struct {
	Method  RequestMethod
	Path    string
	Handler func(c echo.Context) error
}

func (c Controller) GetMethod() RequestMethod {
	return c.Method
}

func (c Controller) GetPath() string {
	return c.Path
}

func (c Controller) GetHandler() func(c echo.Context) error {
	return c.Handler
}
