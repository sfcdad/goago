package htmlbo

import (
	"github.com/goago/app_code/common"
	"fmt"
	"net/http"
	"github.com/goago/app_code/shopapp"
	"github.com/gin-gonic/gin"
)

func Html_ProductComment_Get(c *gin.Context) {
	var productguid = c.DefaultQuery("productguid","0")
	//
	var info shopapp.ProductInfo
	info = shopapp.GetInfo_Product(productguid)

	c.HTML(http.StatusOK, "z_product_comment.html", info)
}

func Html_ProductComment_Post(c *gin.Context) {

	var productguid = c.PostForm("hid_productguid")
	var txt = c.PostForm("txt_productcomment")
	//
	shopapp.Update_Product_ProductComment(productguid,txt)
	//
	var info shopapp.ProductInfo
	info = shopapp.GetInfo_Product(productguid)
	//
	info.StatusText = fmt.Sprintf("Save OK! %s",common.GetBeiJingTime())
	c.HTML(http.StatusOK, "z_product_comment.html", info)
}

func Html_ProductDesText_Get(c *gin.Context) {
	var productguid = c.DefaultQuery("productguid","0")
	//
	var info shopapp.ProductInfo
	info = shopapp.GetInfo_Product(productguid)

	c.HTML(http.StatusOK, "z_product_productdestext.html", info)
}

func Html_ProductDesText_Post(c *gin.Context) {
	var productguid = c.PostForm("hid_productguid")
	var txt = c.PostForm("txt_productdestext")
	//
	shopapp.Update_Product_ProductDesText(productguid,txt)
	//
	var info shopapp.ProductInfo
	info = shopapp.GetInfo_Product(productguid)
	//
	info.StatusText = fmt.Sprintf("Save OK! %s",common.GetBeiJingTime())
	c.HTML(http.StatusOK, "z_product_productdestext.html", info)
}