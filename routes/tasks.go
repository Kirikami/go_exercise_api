package routes

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"time"

	"github.com/kirikami/go_exercise_api/database"
	u "github.com/kirikami/go_exercise_api/utils"
)

func SaveTaskHandler(c echo.Context) error {
	const timeForm = "3 04 PM"
	var id, ids int64

	currentTime := time.Now()
	tasks := []database.Task{}

	db := c.Get("DBConnection").(*gorm.DB)

	err := db.Find(&tasks).Count(&ids).Error

	if err != nil {
		return err
	}

	id = int64(ids + 1)
	task := &database.Task{
		Id: id,
	}

	err = c.Bind(task)

	if err != nil {
		return err
	}

	task.CreatedAt = &currentTime
	task.UpdatedAt = &currentTime
	task.IsDeleted = false
	task.IsCompleted = false

	err = db.Save(&task).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, task)

}

func UpdateTaskHandler(c echo.Context) error {
	const timeForm = "3 04 PM"
	currentTime := time.Now()
	db := c.Value("DBConnection").(*gorm.DB)
	idParam := c.Param("id")

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

	err = c.Bind(task)

	if err != nil {
		return err
	}
	task.UpdatedAt = &currentTime
	if task.IsCompleted == true {
		task.CompletedAt = &currentTime
	}
	//task.IsCompleted, err = strconv.ParseBool(isCompleted)

	err = db.Save(&task).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c echo.Context) error {

	db := c.Value("DBConnection").(*gorm.DB)

	idParam := c.Param("id")

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

func GetTaskHandler(c echo.Context) error {
	db := c.Value("DBConnection").(*gorm.DB)

	idParam := c.Param("id")

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

func GetAllTasksHendler(c echo.Context) error {

	db := c.Value("DBConnection").(*gorm.DB)

	tasks := []database.Task{}
	err := db.Find(&tasks).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}
