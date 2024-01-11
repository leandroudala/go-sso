package helper

import (
	"strconv"
	"udala/sso/exception"
)

func StringToUint64(strId string) (uint64, exception.ApplicationException) {
	id, err := strconv.ParseUint(strId, 10, 64)

	if err != nil {
		return 0, exception.InvalidIDException(strId)
	}

	return id, exception.NilError()
}
