package utils

import (
	"fmt"
	"go-simple-api/internal/delivery/http/dto"
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

func BuildUpdateMap(req *dto.UpdateUserRequest) map[string]interface{} {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	if req.KtpNo != nil {
		updates["ktp_no"] = *req.KtpNo
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.PhoneNumber != nil {
		updates["phone_number"] = *req.PhoneNumber
	}

	return updates
}
