package main

import (
	"github.com/labstack/echo"
)

// is interface for centralized service management
type IMicroservice interface {
	Start() error
	Cleanup() error

	//Http services
	GET(path string, h ServiceHandleFunc)
	POST(path string, h ServiceHandleFunc)
	PUT(path string, h ServiceHandleFunc)
	PATCH(path string, h ServiceHandleFunc)
	DELETE(path string, h ServiceHandleFunc)
}

// is centralized service management
type Microservice struct {
	echo *echo.Echo
}

// is handler for each microservice
type ServiceHandleFunc func(ctx IContext) error

// is the constructor function of microservice
func NewMicroservice() *Microservice {
	return &Microservice{
		echo: echo.New(),
	}
}
