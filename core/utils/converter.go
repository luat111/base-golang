package utils

import (
	"strconv"
)

func ToNumber(t any) int {
	switch t.(type) {
	case int:
		return t.(int)
	case string:
		value, err := strconv.Atoi(t.(string))

		if err != nil {
			return 0
		}

		return value
	default:
		return 0
	}
}
