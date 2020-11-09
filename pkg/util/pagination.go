package util

import (
	"github.com/gin-gonic/gin"
	"github.com/han/go-gin-example/pkg/setting"
	"github.com/unknwon/com"
)

/**
	根据查询条件中的页码获取查询条数
	@param c 当前上下文
	@returns 查询条数
**/
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
