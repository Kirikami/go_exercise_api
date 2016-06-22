package database

import (
	"time"
)

type Task struct {
	Id          int64      `sql:"id" json:"id"`
	Title       string     `sql:"char(255)" json:"title"`
	Description string     `sql:"text" json:"description"`
	Priority    int        `sql:"bool" json:"priority"`
	CreatedAt   *time.Time `sql:"json:"createdAt"`
	UpdatedAt   *time.Time `sql:"datetime"json:"updatedAt"`
	CompletedAt *time.Time `sql:"datetime" json:"completedAt"`
	IsDeleted   bool       `sql:"bool" json:"isDeleted"`
	IsCompleted bool       `sql:"bool" json:"isCompeted"`
}
