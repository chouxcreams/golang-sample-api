package main

import (
	"github.com/labstack/echo"
	"golang-sample-api/infra/prismaRepository"
	"golang-sample-api/prisma/db"
	"golang-sample-api/usecase"
	"log"
	"net/http"
	"strconv"
)

var taskUsecase usecase.TaskUsecase

type TaskCommand struct {
	Text string `json:"text"`
}

type TaskResponse struct {
	text string `json:"text"`
	id   int    `json:"id"`
}

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()
	taskRepository := prismaRepository.NewTaskRepository(client)
	taskUsecase = usecase.NewTaskUsecase(taskRepository)
	e := echo.New()
	e.GET("/", hello)
	e.POST("/tasks", postTasks)
	e.PUT("/tasks/:id", putTask)
	e.GET("/tasks/:id", getTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.Logger.Fatal(e.Start(":8989"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello")
}

func postTasks(c echo.Context) error {
	var task TaskCommand
	if err := c.Bind(&task); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}
	err := taskUsecase.Create(task.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func putTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var task TaskCommand
	if err := c.Bind(&task); err != nil {
		c.Logger().Error("Bind: ", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}
	_, err = taskUsecase.Update(id, task.Text)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	task, err := taskUsecase.FindByID(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, TaskResponse{text: task.Text, id: task.Id})
}

func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = taskUsecase.Delete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
