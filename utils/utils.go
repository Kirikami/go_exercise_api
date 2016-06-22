package utils

import (
	"strconv"
)

func ParseIdInt64FromString(s string) (int64, error) {
	result, err := strconv.ParseInt(c.FormValue("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	return result, nil
}
