package handler

import "strings"

func checkDatabaseError(err error) error {
	msg := err.Error()

	if strings.Contains(msg, "record not found") {
		return NotFoundError("Duplicated Key")
	}

	return UnknownError("Unknown Error")
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	errorVerified := checkDatabaseError(err)
	_, ok := errorVerified.(NotFoundError)

	return ok
}

type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}

type UnknownError string

func (e UnknownError) Error() string {
	return string(e)
}
