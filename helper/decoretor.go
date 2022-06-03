package helper

import (
	"github.com/gin-gonic/gin"
	"mydns/constant"
	"mydns/logs"
	"mydns/utils"
	"net"
	"os"
	"runtime/debug"
	"strings"
)

type OptionalParam struct {
	CheckSign  bool
	CheckUser  bool
	NeedLogApi bool
}

func DefaultOptionalParam() *OptionalParam {
	return &OptionalParam{
		CheckSign:  true,
		CheckUser:  true,
		NeedLogApi: true,
	}
}

func GetLogId(ctx *gin.Context) string {
	logId := ctx.Request.Header.Get("log_id")
	if logId == "" {
		logId = utils.GenerateLogId()
	}

	return logId
}

func Api(f func(*gin.Context) ([]byte, *constant.BaseStatus), dparam *OptionalParam) func(*gin.Context) {
	handler := func(ctx *gin.Context) {
		ctx.Set("uri", ctx.Request.RequestURI)
		ctx.Set("log_id", GetLogId(ctx))

		defer func() {
			if err := recover(); err != nil {
				logs.ErrorCtx(ctx, "[Api] panic recovered:%s, error stack:%s", err, debug.Stack())
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				statusCode := constant.ERR_UNKNOW.StatusCode
				statusMsg := constant.ERR_UNKNOW.StatusMsg
				if brokenPipe {
					statusCode = constant.ERR_USER_CLOSE_CONN.StatusCode
					statusMsg = constant.ERR_USER_CLOSE_CONN.StatusMsg
				}

				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					ctx.Error(err.(error))
					ctx.Abort()
					return
				}
				logs.ErrorCtx(ctx, "[Api] panic error stack:%s", debug.Stack())
				ctx.JSON(200, gin.H{
					"status_code": statusCode,
					"status_msg":  statusMsg,
				})
			}
		}()

		response, status := f(ctx)
		if status != constant.SUCCESS {
			ctx.JSON(200, gin.H{
				"status_code": status.StatusCode,
				"status_msg":  status.StatusMsg,
			})
		} else if response != nil {
			ctx.Data(200, "application/json; charset=utf-8", response)
		}
	}

	return handler
}
