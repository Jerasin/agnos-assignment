package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"agnos-assignment/app/constant"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func PanicDatabaseException(err error, c *gin.Context, skipError error) {
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey) && !errors.Is(err, skipError):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "username or password is exits"))
			return
		case errors.Is(err, gorm.ErrInvalidDB) && !errors.Is(err, skipError):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "invalid database"))
			return
		case errors.Is(err, gorm.ErrInvalidValue) && !errors.Is(err, skipError):
			c.JSON(http.StatusBadRequest, BuildResponse(constant.Duplicated, "invalid value"))
			return
		case errors.Is(err, gorm.ErrRecordNotFound) && !errors.Is(err, skipError):
			c.JSON(http.StatusNotFound, BuildResponse(constant.DataNotFound, "record not found"))
			return
		default:
			if !errors.Is(err, skipError) {
				log.Error("Happened error when saving data to database. Error", err)
				PanicException(constant.UnknownError)
			}

			return
		}
	}
}

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func CustomPanicException(responseKey constant.ResponseStatus, message string) {
	PanicException_(responseKey.GetResponseStatus(), message)
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {

		log.Errorf("PanicHandler = %+v  \n", err)

		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		log.Errorf("key = %+v  \n", key)

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusNotFound, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.BadRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.DataIsExit.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}
