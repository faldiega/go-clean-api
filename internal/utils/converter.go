package utils

import (
	"fmt"
	"strconv"
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
