package routes

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"

	m "github.com/kirikami/go_exercise_api/database/models"
	u "github.com/kirikami/go_exercise_api/utils"
)

func (h ApiV1Handler) SaveTaskHandler(c echo.Context) error {
	task := m.Task{}

	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := h.DB.Save(&task).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, task)
}

func (h ApiV1Handler) UpdateTaskHandler(c echo.Context) error {
	id, err := u.ParseIdInt64FromString(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	task := m.Task{}

	if err := h.DB.First(&task, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := c.Bind(&task); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	task.SetIsCompleted()

	if err := h.DB.Save(&task).Error; err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

func (h ApiV1Handler) DeleteTaskHandler(c echo.Context) error {
	id, err := u.ParseIdInt64FromString(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	task := m.Task{}

	if err := h.DB.First(&task, id).Error; err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)

	}

	task.SetIsDeleted()

	if err := h.DB.Save(&task).Error; err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h ApiV1Handler) GetTaskHandler(c echo.Context) error {
	id, err := u.ParseIdInt64FromString(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	task := m.Task{}

	if err := h.DB.First(&task, id).Error; err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, task)
}

func (h ApiV1Handler) GetAllTasksHendler(c echo.Context) error {
	tasks := m.TaskList{}

	if err := h.DB.Find(&tasks.Tasks).Error; err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, tasks)
}
