package task

import (
	"context"
	"try_go_swagger/common"
	sqlcDB "try_go_swagger/gen/db"
	"try_go_swagger/gen/models"
	taskOperations "try_go_swagger/gen/restapi/operations/task"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func GetTask(p taskOperations.GetTaskParams) middleware.Responder {
	db, err := common.CreateDB()
	if err != nil {
		return taskOperations.NewGetTaskInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}
	defer db.Close()

	query := sqlcDB.New(db)
	taskList, err := query.GetAllTaskList(context.Background())
	if err != nil {
		return taskOperations.NewGetTaskInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}

	var resultTaskList []*models.Task
	for _, task := range taskList {
		formatedTask := models.Task{
			TaskID:      int64(task.Taskid),
			Status:      &task.Status.String,
			Title:       task.Title.String,
			Desc:        task.Desc.String,
			CreatedDate: strfmt.DateTime(task.CreatedDate.Time),
		}
		resultTaskList = append(resultTaskList, &formatedTask)
	}

	return taskOperations.NewGetTaskOK().WithPayload(resultTaskList)
}
