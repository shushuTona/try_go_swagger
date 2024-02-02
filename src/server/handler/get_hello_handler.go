package handler

import (
	"try_go_swagger/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetHello(p operations.HelloParams) middleware.Responder {
	return operations.NewHelloOK().WithPayload("Hello Swagger")
}
