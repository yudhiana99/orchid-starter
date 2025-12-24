package constants

import (
	"strconv"

	"github.com/mataharibiz/sange/v2"
)

func GetConstant(statusName string) int64 {
	statusConstant, errGet := sange.GetConstant(statusName)
	if errGet != nil {
		return 0
	}

	status, errParse := strconv.ParseInt(statusConstant, 10, 64)
	if errParse != nil {
		return 0
	}

	return status
}
