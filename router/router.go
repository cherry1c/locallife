package router

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/business/homepage"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/slides", homepage.GetSwiperList)
	router.POST("/image", homepage.AddImageRes)

	return router
}
