package hello

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mydns/constant"
)

func Hello(ctx *gin.Context) ([]byte, *constant.BaseStatus) {

	return []byte("hello"), constant.SUCCESS
}

func TestPanic(ctx *gin.Context) ([]byte, *constant.BaseStatus) {

	panic(errors.New("test"))

	return []byte("hello"), constant.SUCCESS
}

func TestError(ctx *gin.Context) ([]byte, *constant.BaseStatus) {

	return nil, constant.ERR_INVALID_PARAM
}
