package utils

import (
	"strconv"
)

var result int64

func ParseIdInt64FromString(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 10, 64)

	return result, err
}
