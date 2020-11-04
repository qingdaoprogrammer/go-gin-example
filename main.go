package main

import (
	"fmt"
	"net/http"

	"github.com/han/go-gin-example/pkg/setting"
	"github.com/han/go-gin-example/routers"
)

func main() {
	/*router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})*/
	router := routers.InitRouter() //初始化路由
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
