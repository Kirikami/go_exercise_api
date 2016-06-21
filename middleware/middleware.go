package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
	"github.com/labstack/echo"
)

func UseConfig(db *gorm.DB, config *config.Configuration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			c.Set("config", config)
			return next(c)
		}
	}
}
