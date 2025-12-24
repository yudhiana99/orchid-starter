package common

import (
	"errors"
	"strings"

	"github.com/mataharibiz/sange/v2"
)

func GetChainError(err error) (msg string) {
	chainString := " -> "
	for err != nil {
		msg += err.Error() + chainString
		err = errors.Unwrap(err)
	}
	return strings.TrimSuffix(msg, chainString)
}

func IsSangeError(err error) (ok bool, errSange *sange.AppError) {
	if err != nil {
		if errSange, ok = err.(*sange.AppError); ok {
			return
		}
	}

	return false, nil
}
