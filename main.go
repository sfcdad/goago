/*
jing > ping > taobao
2018-8-2
*/
package main

import (
	// "github.com/goago/app_code/webms"
	"fmt"
	
	"github.com/gin-gonic/gin"
)

func Init() {
	
	fmt.Println("init done.")
}

func main() {
	// webms.InsertTestData_Menu()
	fmt.Println("start Port:8080..")
	r := gin.Default()
	//
	allRouter(r, fmt.Sprintf("%s", "/"))
	allStatic(r) //静态文件
	allTemplates(r)
	//
	r.Run()
}

func allRouter(router *gin.Engine, prefix string) {
	RouterBO(router, prefix)
}

func allStatic(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/dbimg", "./dbimg")
}

func allTemplates(router *gin.Engine) {

	router.LoadHTMLGlob("html/*")

}
