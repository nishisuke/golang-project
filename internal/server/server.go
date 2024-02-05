package server

import (
	"golang-project/internal/handler"
	"golang-project/internal/repo"
	"golang-project/internal/usecase/taskusecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) (*echo.Echo, error) {
	repo := repo.NewTaskRepo(db)

	taskListUsecase := taskusecase.NewListUsecase(repo)
	taskCreateUsecase := taskusecase.NewCreateUsecase(repo)
	taskToggleDoneUsecase := taskusecase.NewToggleDoneUsecase(repo)
	taskDeleteUsecase := taskusecase.NewDeleteUsecase(repo)

	taskListHandler := handler.NewTaskListHandler(taskListUsecase)
	taskCreateHandler := handler.NewTaskCreateHandler(taskCreateUsecase)
	taskToggleDoneHandler := handler.NewTaskToggleDoneHandler(taskToggleDoneUsecase)
	taskDeleteHandler := handler.NewTaskDeleteHandler(taskDeleteUsecase)

	e := echo.New()
	e.GET("/tasks", taskListHandler.EchoHandler)
	e.POST("/tasks", taskCreateHandler.EchoHandler)
	e.PATCH("/tasks/:id", taskToggleDoneHandler.EchoHandler)
	e.DELETE("/tasks/:id", taskDeleteHandler.EchoHandler)

	return e, nil
}
