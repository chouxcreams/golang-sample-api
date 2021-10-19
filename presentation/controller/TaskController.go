package controller

import (
	"github.com/labstack/echo"
	"golang-sample-api/presentation/apiModel"
	"golang-sample-api/usecase"
	"net/http"
	"strconv"
)

type TaskController interface {
	GetTask(c echo.Context) error
	PostTasks(c echo.Context) error
	PutTask(c echo.Context) error
	DeleteTask(c echo.Context) error
	CompleteTask(c echo.Context) error
	GetTaskList(c echo.Context) error
}

type taskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(taskUsecase usecase.TaskUsecase) TaskController {
	return &taskController{
		taskUsecase: taskUsecase,
	}
}

func (tc *taskController) PostTasks(c echo.Context) error {
	var task apiModel.TaskCommand
	if err := c.Bind(&task); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}
	err := tc.taskUsecase.Create(task.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (tc *taskController) PutTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var task apiModel.TaskCommand
	if err := c.Bind(&task); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}
	_, err = tc.taskUsecase.Update(id, task.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	task, err := tc.taskUsecase.FindByID(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, apiModel.NewTaskResponse(task))
}

func (tc *taskController) CompleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = tc.taskUsecase.Complete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = tc.taskUsecase.Delete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) GetTaskList(c echo.Context) error {
	tasks, err := tc.taskUsecase.FindAll()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var taskResponses []apiModel.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, apiModel.NewTaskResponse(&task))
	}
	return c.JSON(http.StatusOK, taskResponses)
}
