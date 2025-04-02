package helpers

import (
	"github.com/alaa-aqeel/govalid/src/pkgs/database"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(message string, data any) gin.H {
	return gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(message string, errors any) gin.H {
	return gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	}
}

func DatabaseErrorResponse(err error) gin.H {
	sqlError := database.HandleSqlError(err)

	return gin.H{
		"status":  "error",
		"message": sqlError.Error(),
	}
}
