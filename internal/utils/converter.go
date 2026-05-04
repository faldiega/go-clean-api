package utils

import (
	"fmt"
	"strconv"
	"time"
)

func StringToInt(s string) int {

	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("String to Int conversion error:")
		panic(err)
	}

	return result
}

func StringToBool(s string) bool {

	result, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Println("String to Bool conversion error:")
		panic(err)
	}

	return result
}

// parse string ke int positif, return default jika invalid
func ParsePositiveInt(s string, defaultVal int) int {
	val, err := strconv.Atoi(s)
	if err != nil || val < 1 {
		return defaultVal
	}
	return val
}

func ToDatetime(date any, args ...string) *string {
	format := time.DateTime // nilai default
	var result time.Time

	switch val := date.(type) {
	case time.Time:
		result = val
	case *time.Time:
		if val == nil {
			return nil
		}
		result = *val
	default:
		return nil
	}

	if len(args) > 0 {
		format = args[0]
	}

	formatted := result.Format(format)
	return &formatted
}
