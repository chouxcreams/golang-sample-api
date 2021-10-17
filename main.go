package main

import (
	"github.com/labstack/echo"
	"golang-sample-api/infra/prismaRepository"
	"golang-sample-api/presentation/controller"
	"golang-sample-api/prisma/db"
	"golang-sample-api/usecase"
	"log"
	"net/http"
)

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
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)
	e := echo.New()
	e.GET("/", hello)
	e.POST("/tasks", taskController.PostTasks)
	e.GET("/tasks/:id", taskController.GetTask)
	e.PUT("/tasks/:id", taskController.PutTask)
	e.PUT("/tasks/:id/complete", taskController.CompleteTask)
	e.DELETE("/tasks/:id", taskController.DeleteTask)
	e.Logger.Fatal(e.Start(":8989"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello")
}
