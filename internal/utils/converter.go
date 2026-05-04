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

// format datetime dd-MM-yyyy hh:mm:ss
func ToDatetime(date time.Time) string {
	return date.Format(time.DateTime)
}

// format datetime dd-MM-yyyy hh:mm:ss untuk nullable value
func ToDatetimeNullable(date *time.Time) *string {
	if date == nil {
		return nil
	}

	formatted := date.Format(time.DateTime)
	return &formatted
}
