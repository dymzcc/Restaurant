package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 1 //操作成功
	FAILED  int = 0 //操作失败
)

//普通操作成功返回
func Success(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})
}

//普通操作失败返回
func Failed(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"smg":  v,
	})
}
