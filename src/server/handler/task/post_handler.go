package task

import (
	"context"
	"database/sql"
	"try_go_swagger/common"
	sqlcDB "try_go_swagger/gen/db"
	"try_go_swagger/gen/models"
	taskOperations "try_go_swagger/gen/restapi/operations/task"

	"github.com/go-openapi/runtime/middleware"
)

func AddTask(p taskOperations.AddTaskParams) middleware.Responder {
	db, err := common.CreateDB()
	if err != nil {
		return taskOperations.NewGetTaskInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}
	defer db.Close()

	query := sqlcDB.New(db)
	arg := sqlcDB.CreateTaskParams{
		Status: sql.NullString{String: "Waiting", Valid: true},
		Title:  sql.NullString{String: p.Body.Title, Valid: true},
		Desc:   sql.NullString{String: p.Body.Desc, Valid: true},
	}
	_, err = query.CreateTask(context.Background(), arg)
	if err != nil {
		return taskOperations.NewGetTaskInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}

	return taskOperations.NewAddTaskOK()
}
