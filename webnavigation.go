package main

import (
	"github.com/goago/app_code/htmlbo"
	"github.com/goago/app_code/shopapp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterBO(router *gin.Engine, prefix string) {

	//home_managementstudio2018
	ms := fmt.Sprintf("%s/%s", prefix, "ms")
	router.GET(ms, htmlbo.Html_MS2018)

	// 首页
	index := fmt.Sprintf("%s/%s", prefix, "")
	router.GET(index, Index)

	// 登录
	// login := fmt.Sprintf("%s/%s", prefix, "login")
	// router.GET(login, LoginHTML)
	// router.POST(login, Login)

	//producthome
	product := fmt.Sprintf("%s/%s", prefix, "product")
	router.GET(product, htmlbo.Html_ShowProduct)
	router.POST(product, htmlbo.Html_ShowProduct_ToPage)

	//productedit
	productinfo := fmt.Sprintf("%s/%s", prefix, "productinfo")
	router.GET(productinfo, htmlbo.Html_ShowProductInfo)
	router.POST(productinfo, htmlbo.Html_ProductEdit)

	//product_add
	product_add := fmt.Sprintf("%s/%s", prefix, "product_add")
	router.GET(product_add, htmlbo.Html_ProductAddNew)
	router.POST(product_add, htmlbo.Html_ProductAddNewSave)

	//product_delete
	product_delete := fmt.Sprintf("%s/%s", prefix, "product_delete")
	router.POST(product_delete, htmlbo.Html_ProductDelete)

	//ProductComment
	productcomment := fmt.Sprintf("%s/%s", prefix, "productcomment")
	router.GET(productcomment, htmlbo.Html_ProductComment_Get)
	router.POST(productcomment, htmlbo.Html_ProductComment_Post)

	//ProductDesText
	productdestext := fmt.Sprintf("%s/%s", prefix, "productdestext")
	router.GET(productdestext, htmlbo.Html_ProductDesText_Get)
	router.POST(productdestext, htmlbo.Html_ProductDesText_Post)

	//explorerlist
	explorerlist := fmt.Sprintf("%s/%s", prefix, "explorerlist")
	router.GET(explorerlist, htmlbo.Html_ExplorerList)
	router.POST(explorerlist, htmlbo.Html_ExplorerList_Update)

	//explorermenu
	explorermenu := fmt.Sprintf("%s/%s", prefix, "explorermenu")
	router.GET(explorermenu, htmlbo.Html_ExplorerMenu)

	//explorerlistadd
	explorerlistadd := fmt.Sprintf("%s/%s", prefix, "explorerlistadd")
	router.GET(explorerlistadd, htmlbo.Html_ExplorerList_AddNew)

	//explorersavetoplevel 保存大的分类
	explorersavetoplevel := fmt.Sprintf("%s/%s", prefix, "explorersavetoplevel")
	router.POST(explorersavetoplevel, htmlbo.Html_Update_ExplorerTopLevel)

	//Insert_ProductExplorer
	insert_productExplorer := fmt.Sprintf("%s/%s", prefix, "insert_productExplorer")
	router.GET(insert_productExplorer, htmlbo.Html_InsertProductExplorer)

	productlabelset := fmt.Sprintf("%s/%s", prefix, "productlabelset")
	router.GET(productlabelset, htmlbo.Html_ProductLabelSet)
	router.POST(productlabelset, htmlbo.Html_ProductLabelSet_Save)
}

func Index(c *gin.Context) {

	var info shopapp.HomePageInfo = shopapp.GetInfo_HomePage()
	//
	c.HTML(http.StatusOK, "start.html",info)
}
