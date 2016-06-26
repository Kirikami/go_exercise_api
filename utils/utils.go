package utils

import (
	"github.com/kirikami/go_exercise_api/database"
	"strconv"
	"time"
)

func ParseIdInt64FromString(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)

	return result, err
}

func SetIsCompleted(t *database.Task) {
	current_time := time.Now()
	if t.IsCompleted == true {
		t.CompletedAt = &current_time
	}
}
