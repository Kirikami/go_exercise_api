package routes

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"

	"github.com/kirikami/go_exercise_api/database"
	u "github.com/kirikami/go_exercise_api/utils"
)

func (h ApiV1Handler) SaveTaskHandler(c echo.Context) error {
	task := database.Task{}

	err := c.Bind(&task)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = h.DB.Save(&task).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, task)
}

func (h ApiV1Handler) UpdateTaskHandler(c echo.Context) error {
	idParam := c.P(0)

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	task := database.Task{}
	err = h.DB.First(&task, id).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = c.Bind(&task)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	u.SetIsCompleted(&task)

	err = h.DB.Save(&task).Error

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

func (h ApiV1Handler) DeleteTaskHandler(c echo.Context) error {
	idParam := c.P(0)

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	task := database.Task{}
	err = h.DB.First(&task, id).Error

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)

	}

	task.IsDeleted = true
	err = h.DB.Save(&task).Error

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h ApiV1Handler) GetTaskHandler(c echo.Context) error {
	idParam := c.P(0)

	id, err := u.ParseIdInt64FromString(idParam)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	task := database.Task{}
	err = h.DB.First(&task, id).Error

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

func (h ApiV1Handler) GetAllTasksHendler(c echo.Context) error {
	tasks := []database.Task{}
	err := h.DB.Find(&tasks).Error

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, tasks)
}
