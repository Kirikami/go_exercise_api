package routes

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"time"

	"github.com/kirikami/go_exercise_api/database"
	u "github.com/kirikami/go_exercise_api/utils"
)

func (h ApiV1Handler) SaveTaskHandler(c echo.Context) error {
	task := database.Task{}

	db := h.DB

	err := c.Bind(&task)

	if err != nil {
		return err
	}

	err = db.Save(&task).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, task)

}

func (h ApiV1Handler) UpdateTaskHandler(c echo.Context) error {
	currentTime := time.Now()
	db := h.DB
	idParam := c.P(0)
	fmt.Println(idParam)

	if idParam == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return err
	}

	task := database.Task{}
	err = db.First(&task, id).Error

	if err != nil {
		return err
	}

	err = c.Bind(&task)

	if err != nil {
		return err
	}
	if task.IsCompleted == true {
		task.CompletedAt = &currentTime
	}

	err = db.Save(&task).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (h ApiV1Handler) DeleteTaskHandler(c echo.Context) error {
	db := h.DB

	idParam := c.P(0)

	if idParam == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return err
	}

	task := database.Task{}
	err = db.First(&task, id).Error

	if err != nil {
		return err
	}

	task.IsDeleted = true
	err = db.Save(&task).Error

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h ApiV1Handler) GetTaskHandler(c echo.Context) error {
	db := h.DB

	idParam := c.P(0)
	fmt.Println(idParam)

	if idParam == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return err
	}

	task := database.Task{}
	err = db.First(&task, id).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)

}

func (h ApiV1Handler) GetAllTasksHendler(c echo.Context) error {
	db := h.DB

	tasks := []database.Task{}
	err := db.Find(&tasks).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}
