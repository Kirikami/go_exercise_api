package routes

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"

	"github.com/kirikami/go_exercise_api/database"
)

func SaveTaskHandler(c echo.Context) error {
	const timeForm = "3 04 PM"
	var ids int
	var id int64

	currentTime := time.Now()
	tasks := []database.Task{}

	db := c.Get("DBConnection").(*gorm.DB)

	if c.FormValue("id") != "" {
		id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		_ = id

	} else {
		err := db.Find(&tasks).Count(&ids).Error
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		id = int64(ids + 1)
	}

	title := c.FormValue("title")
	description := c.FormValue("description")
	priority, err := strconv.Atoi(c.FormValue("priority"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	createdAt := &currentTime
	updatedAt := &currentTime
	isDeleted := false
	isCompeted := false

	taskDocument := database.Task{id, title, description, priority, createdAt, updatedAt, nil, isDeleted, isCompeted}

	err = db.Save(&taskDocument).Error
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(302, "/")
}

func UpdateTaskHandler(c echo.Context) error {
	const timeForm = "3 04 PM"
	currentTime := time.Now()
	db := c.Value("DBConnection").(*gorm.DB)

	id := c.Param("id")
	if id == "" {
		return c.Redirect(302, "/")
	}

	task := database.Task{}
	err := db.First(&task, id).Error
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	task.Title = c.FormValue("title")
	task.Description = c.FormValue("description")
	task.Priority, err = strconv.Atoi(c.FormValue("priority"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	task.UpdatedAt = &currentTime
	if c.FormValue("isCompeted") == "true" {
		task.CompletedAt = &currentTime
	}
	isCompleted := c.FormValue("isCompleted")
	task.IsCompleted, err = strconv.ParseBool(isCompleted)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	err = db.Save(&task).Error

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "write", task)
}

func DeleteTaskHandler(c echo.Context) error {

	db := c.Value("DBConnection").(*gorm.DB)

	id := c.Param("id")
	if id == "" {
		return c.Redirect(302, "/")
	}

	task := database.Task{}
	err := db.First(&task, id).Error

	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	task.IsDeleted = true
	err = db.Save(&task).Error

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(302, "/")
}

func WriteTaskHandler(c echo.Context) error {
	task := database.Task{}
	return c.Render(http.StatusOK, "write", task)
}

func GetTaskHandler(c echo.Context) error {

	db := c.Value("DBConnection").(*gorm.DB)

	id := c.Param("id")

	if id == "" {
		return c.Redirect(302, "/")
	}

	task := database.Task{}
	err := db.First(&task, id).Error

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "view", task)
}

func GetAllTasksHendler(c echo.Context) error {

	db := c.Value("DBConnection").(*gorm.DB)

	tasks := []database.Task{}
	err := db.Find(&tasks).Error

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Render(http.StatusOK, "index", tasks)
}
