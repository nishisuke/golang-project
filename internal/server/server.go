package server

import (
	"errors"
	"golang-project/internal/handler"
	"golang-project/internal/pkg/mailer"
	"golang-project/internal/repo"
	"golang-project/internal/usecase/taskusecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) (*echo.Echo, error) {
	repo := repo.NewTaskRepo(db)

	mailer := mailer.NewMailer()

	taskCreateService := taskusecase.NewCreateTaskService(db, mailer)

	taskListUsecase := taskusecase.NewListUsecase(repo)
	taskCreateUsecase := taskusecase.NewCreateUsecase(repo, taskCreateService)
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

	orgHandler := e.HTTPErrorHandler
	e.HTTPErrorHandler = customHTTPErrorHandler(orgHandler)

	return e, nil
}

func json(c echo.Context, code int, i interface{}) {
	if err := c.JSON(code, i); err != nil {
		c.Logger().Error(err)
	}
}

type (
	ValidationErrorResponse struct {
		ValidationError ValidationError `json:"validation_error"`
	}
	ValidationError struct {
		Message string                 `json:"message"`
		Fields  []ValidationErrorField `json:"fields"`
	}

	ValidationErrorField struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func customHTTPErrorHandler(fallback func(error, echo.Context)) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var response ValidationErrorResponse
			for _, fieldError := range validationErrors {
				response.ValidationError.Fields = append(response.ValidationError.Fields, ValidationErrorField{
					Field:   fieldError.Field(),
					Message: fieldError.Error(),
				})
			}

			json(c, http.StatusBadRequest, response)
			return
		}
		fallback(err, c)
	}
}
