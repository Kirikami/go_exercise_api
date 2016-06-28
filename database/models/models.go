package models

import (
	"time"
)

type Task struct {
	Id          int64      `sql:"id" json:"id"`
	Title       string     `sql:"title" json:"title"`
	Description string     `sql:"description" json:"description"`
	Priority    int        `sql:"priority" json:"priority"`
	CreatedAt   time.Time  `sql:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `sql:"updated_at" json:"updatedAt"`
	CompletedAt *time.Time `sql:"completed_at" json:"completedAt"`
	IsDeleted   bool       `sql:"is_deleted" json:"isDeleted"`
	IsCompleted bool       `sql:"is_competed" json:"isCompeted"`
}

func (t Task) SetIsCompleted() {
	current_time := time.Now()
	if t.IsCompleted == true {
		t.CompletedAt = &current_time
	}
}
