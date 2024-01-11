package exception

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApplicationException struct {
	StatusCode uint
	Message    string
}

func BadRequestException(message string) ApplicationException {
	return ApplicationException{
		StatusCode: 400,
		Message:    message,
	}
}

func (err ApplicationException) HasError() bool {
	return err.StatusCode != 0
}

func (err ApplicationException) Abort(c *gin.Context) {
	c.AbortWithStatusJSON(
		int(err.StatusCode),
		gin.H{
			"error": err.Message,
		},
	)
}

func NilError() ApplicationException {
	return ApplicationException{}
}

func UserNotFoundException(id uint64) ApplicationException {
	return ApplicationException{
		StatusCode: 404,
		Message:    "User id '" + strconv.FormatUint(id, 10) + "' not found",
	}
}

func UserDuplicatedException() ApplicationException {
	return ApplicationException{
		StatusCode: 409,
		Message:    "Username or email is already in use.",
	}
}

func InternalServerError(message string) ApplicationException {
	return ApplicationException{
		StatusCode: 500,
		Message:    message,
	}
}
