package main

import (
	"github.com/goago/app_code/shopapp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterBO(router *gin.Engine, prefix string) {


	// 首页
	index := fmt.Sprintf("%s/%s", prefix, "")
	router.GET(index, Index)

	// 登录
	// login := fmt.Sprintf("%s/%s", prefix, "login")
	// router.GET(login, LoginHTML)
	// router.POST(login, Login)

}

func Index(c *gin.Context) {

	var info shopapp.HomePageInfo = shopapp.GetInfo_HomePage()
	//
	c.HTML(http.StatusOK, "start.html",info)
}
