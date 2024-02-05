package task

import (
	"fmt"
	taskOperations "try_go_swagger/gen/restapi/operations/task"

	"github.com/go-openapi/runtime/middleware"
)

func AddTask(p taskOperations.AddTaskParams) middleware.Responder {
	fmt.Printf("HTTPRequest : %#v\n", p.HTTPRequest)
	fmt.Printf("Title : %#v\n", p.Body.Title)
	fmt.Printf("Desc : %#v\n", p.Body.Desc)

	return taskOperations.NewAddTaskOK()
}
