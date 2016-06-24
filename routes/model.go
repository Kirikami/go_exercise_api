package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kirikami/go_exercise_api/config"
)

type ApiV1Handler struct {
	DB     *gorm.DB
	Config *config.Configuration
}
