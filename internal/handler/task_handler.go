package handler

import (
	"golang-project/internal/usecase/taskusecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	TaskListHandler struct {
		listUsecase *taskusecase.ListUsecase
	}
	TaskCreateHandler struct {
		createUsecase *taskusecase.CreateUsecase
	}
	TaskToggleDoneHandler struct {
		toggleDoneUsecase *taskusecase.ToggleDoneUsecase
	}
	TaskDeleteHandler struct {
		deleteUsecase *taskusecase.DeleteUsecase
	}
)

func NewTaskListHandler(listUsecase *taskusecase.ListUsecase) *TaskListHandler {
	return &TaskListHandler{listUsecase}
}

func NewTaskCreateHandler(createUsecase *taskusecase.CreateUsecase) *TaskCreateHandler {
	return &TaskCreateHandler{createUsecase}
}

func NewTaskToggleDoneHandler(toggleDoneUsecase *taskusecase.ToggleDoneUsecase) *TaskToggleDoneHandler {
	return &TaskToggleDoneHandler{toggleDoneUsecase}
}

func NewTaskDeleteHandler(deleteUsecase *taskusecase.DeleteUsecase) *TaskDeleteHandler {
	return &TaskDeleteHandler{deleteUsecase}
}

// 下記のようにやるのがフレームワークの影響を吸収できてていねいだが、echoのHandlerFunc型はerrorを返すインターフェイスが優れているのでそのまま使う。
// func (h TaskListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
// e.GET("/tasks", echo.WrapHandler(http.HandlerFunc(taskusecase.NewTaskListHandler(...))))
func (h TaskListHandler) EchoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	input := &taskusecase.ListInput{
		Type: taskusecase.NewTaskListType(c.QueryParam("type")),
	}
	list, err := h.listUsecase.TaskList(ctx, input)
	if err != nil {
		return err
	}
	c.JSONPretty(http.StatusOK, list, "  ")
	return nil
}

func (h TaskCreateHandler) EchoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	input := new(taskusecase.CreateInput)

	if err := c.Bind(input); err != nil {
		return err
	}

	task, err := h.createUsecase.CreateTask(ctx, input)
	if err != nil {
		return err
	}
	c.JSONPretty(http.StatusCreated, task, "  ")
	return nil
}

func (h TaskToggleDoneHandler) EchoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	input := &taskusecase.ToggleDoneInput{}
	err := h.toggleDoneUsecase.ToggleTaskDone(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (h TaskDeleteHandler) EchoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	input := &taskusecase.DeleteInput{}
	err := h.deleteUsecase.DeleteTask(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
