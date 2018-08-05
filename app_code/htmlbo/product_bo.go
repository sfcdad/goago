package htmlbo

import (
	"strconv"
	"github.com/goago/app_code/webms"
	"fmt"
	"bytes"
	"net/http"
	"github.com/goago/app_code/common"
	"strings"
	"github.com/goago/app_code/shopapp"
	"github.com/gin-gonic/gin"
)

var m_msg shopapp.PageProductEditInfo
var m_lbldict = make(map[string]shopapp.ProductLabelInfo)

func Html_ProductLabelSet(c *gin.Context) {

	var cmd = c.DefaultQuery("cmd", "0")
	var productguid = c.DefaultQuery("id", "0")
	var expid = c.DefaultQuery("expid", "0")
	var status = c.DefaultQuery("status", "0")
	//
	if cmd == "1" {
		var msg shopapp.PageProductEditInfo
		msg.PageTitle = "Product Label"
		msg.RowInfo = shopapp.GetInfo_Product(productguid)
		msg.LabelList = shopapp.GetListAll_ProductExplorer()
		m_msg = msg
		//default data
		// var a=msg.RowInfo.ExplorerGUID
		var a = "11,13,15,"
		a = msg.RowInfo.ExplorerGUID
		arr := strings.Split(a, ",")
		var dictgoodslbs = make(map[string]string)
		for kp := range arr {
			var key = arr[kp]
			//
			if key != "" {
				dictgoodslbs[key] = key
			}
		}
		//
		for k := range m_msg.LabelList {

			for k2 := range m_msg.LabelList[k].SecondLevelTagList {

				m_msg.LabelList[k].SecondLevelTagList[k2].ProductGUID = m_msg.RowInfo.ProductGUID
				//
				var LabelGUID = m_msg.LabelList[k].SecondLevelTagList[k2].LabelGUID
				if _, oklbl := dictgoodslbs[LabelGUID]; oklbl {
					//存在
					m_msg.LabelList[k].SecondLevelTagList[k2].ClassStyle = "a_btn_labelselect"
					m_msg.LabelList[k].SecondLevelTagList[k2].Status = 1
				} else {
					m_msg.LabelList[k].SecondLevelTagList[k2].ClassStyle = "a_btn_labeldefault"
					m_msg.LabelList[k].SecondLevelTagList[k2].Status = 0
				}
				//
				m_lbldict[m_msg.LabelList[k].SecondLevelTagList[k2].LabelGUID] = m_msg.LabelList[k].SecondLevelTagList[k2]
			} //level 2
		} //level 1

		for kp := range arr {
			var key = arr[kp]
			if _, ok := m_lbldict[key]; ok {
				//存在
				var info = m_lbldict[key]
				info.ClassStyle = "a_btn_labelselect"
				info.Status = 1
				//
				m_lbldict[key] = info
			}
		}
	}

	if cmd == "2" {
		for k := range m_msg.LabelList {
			//
			for k2 := range m_msg.LabelList[k].SecondLevelTagList {
				if expid == m_msg.LabelList[k].SecondLevelTagList[k2].LabelGUID {
					if status == "0" {
						m_msg.LabelList[k].SecondLevelTagList[k2].Status = 1
						m_msg.LabelList[k].SecondLevelTagList[k2].ClassStyle = "a_btn_labelselect"
						//更新值
						m_lbldict[m_msg.LabelList[k].SecondLevelTagList[k2].LabelGUID] = m_msg.LabelList[k].SecondLevelTagList[k2]

					} else {
						m_msg.LabelList[k].SecondLevelTagList[k2].Status = 0
						m_msg.LabelList[k].SecondLevelTagList[k2].ClassStyle = "a_btn_labeldefault"
						//
						m_lbldict[m_msg.LabelList[k].SecondLevelTagList[k2].LabelGUID] = m_msg.LabelList[k].SecondLevelTagList[k2]
					}
				}
			} //level 2
		} //level 1

	}
	//
	var txt string
	var txtbuf bytes.Buffer
	var txtbufids bytes.Buffer

	for k, v := range m_lbldict {
		// fmt.Printf("Key: %s  Value: %d\n", k, v)
		if v.Status > 0 {
			txtbuf.WriteString(v.LabelName)
			txtbuf.WriteString(",")
			//
			txtbufids.WriteString(k)
			txtbufids.WriteString(",")
		}
	}
	txt = txtbuf.String()
	m_msg.CurrentLabelList = txt
	//
	if cmd == "3" {
		shopapp.Update_ProductSearchLabel(m_msg.RowInfo.ProductGUID, txtbufids.String(), m_msg.CurrentLabelList)
		m_msg.RowInfo.StatusText = fmt.Sprintf("已经成功保存! %s", common.GetBeiJingTime())
	}
	c.HTML(http.StatusOK, "z_product_label.html", m_msg)

}

func Html_ProductLabelSet_Save(c *gin.Context) {
	var productguid = c.DefaultQuery("id", "0")
	// var expid = c.DefaultQuery("expid", "0")
	// var status = c.DefaultQuery("status", "0")
	//

	//
	var msg shopapp.PageProductEditInfo
	msg.PageTitle = "Product Label"
	msg.RowInfo = shopapp.GetInfo_Product(productguid)
	msg.LabelList = shopapp.GetListAll_ProductExplorer()

	c.HTML(http.StatusOK, "z_product_label.html", msg)

}

func Html_ProductAddNew(c *gin.Context) {

	var msg shopapp.PageProductInfo
	msg.PageTitle = "Product Add"
	//
	msg.RowInfo.ProductGUID = common.GetGuid()
	msg.RowInfo.ProductName = "新产品名称"
	msg.RowInfo.IconUrl = "https://img13.360buyimg.com/n1/s450x450_jfs/t16990/124/1933004112/96067/8d773cb/5addb475N8bf4c6ea.jpg"
	msg.RowInfo.SellUrl = "https://img13.360buyimg.com/n1/s450x450_jfs/t16990/124/1933004112/96067/8d773cb/5addb475N8bf4c6ea.jpg"
	msg.RowInfo.OPTime = common.GetNowTimestamp()
	msg.RowInfo.OPTimeText = common.GetBeiJingTime()
	//

	shopapp.Insert_ProductEmptyRow(msg.RowInfo)
	// Html_ShowProduct(c)
	//
	//打开原网页链接
	var durl = "/product"
	c.Redirect(http.StatusFound, durl)

}

func Html_ProductDelete(c *gin.Context) {
	var productguid = c.DefaultQuery("id", "0")
	var delOK = c.PostForm("txt_productname_del")

	var msg shopapp.PageProductInfo
	msg.RowInfo = shopapp.GetInfo_Product(productguid)

	if msg.RowInfo.ProductName == delOK {
		shopapp.Delete_Product(msg.RowInfo)
		//打开原网页链接
		var durl = "/product"
		c.Redirect(http.StatusFound, durl)
	} else {
		c.HTML(http.StatusOK, "z_product_detail.html", msg.RowInfo)
	}
}

func Html_ProductAddNewSave(c *gin.Context) {
	// var msg shopapp.PageProductInfo
	// msg.PageTitle = "Product Add"
	// msg.RowInfo = shopapp.ProductInfo{}
	// //
	// msg.RowInfo.ProductName = c.PostForm("txt_productname")
	// msg.RowInfo.IconUrl = c.PostForm("txt_iconurl")
	// msg.RowInfo.SellUrl = c.PostForm("txt_sellurl")
	// msg.RowInfo.OPTime = common.GetNowTimestamp()
	// msg.RowInfo.OPTimeText = common.GetBeiJingTime()
	// //
	// shopapp.Insert_Product(msg.RowInfo)

	// Html_ShowProduct(c)
}

func Html_ProductEdit(c *gin.Context) {
	var productguid = c.DefaultQuery("id", "0")
	//
	var msg shopapp.PageProductEditInfo
	msg.PageTitle = "Product Detail"
	msg.RowInfo = shopapp.GetInfo_Product(productguid)
	//
	msg.RowInfo.ProductName = c.PostForm("txt_productname")
	msg.RowInfo.IconUrl = c.PostForm("txt_iconurl")
	msg.RowInfo.SellUrl = c.PostForm("txt_sellurl")
	msg.RowInfo.OPTime = common.GetNowTimestamp()
	msg.RowInfo.OPTimeText = common.GetBeiJingTime()
	//接收到的标签集合

	msg.LabelList = shopapp.GetListAll_ProductExplorer()
	for k := range msg.LabelList {
		var chkname = fmt.Sprintf("chk_%s", msg.LabelList[k].ExplorerGUID)
		var chkvalue = c.PostForm(chkname)
		fmt.Println(chkvalue)
	}

	var suc = shopapp.Update_Product(msg.RowInfo)
	if suc == "suc" {
		msg.RowInfo.StatusText = fmt.Sprintf("%s  %s", "已经成功保存!", msg.RowInfo.OPTimeText)
	}
	c.HTML(http.StatusOK, "z_product_detail.html", msg)
}

func Html_ShowProductInfo(c *gin.Context) {
	var productguid = c.DefaultQuery("id", "0")
	//
	var msg shopapp.PageProductEditInfo
	msg.PageTitle = "Product Home"
	msg.RowInfo = shopapp.GetInfo_Product(productguid)
	msg.LabelList = shopapp.GetListAll_ProductExplorer()

	c.HTML(http.StatusOK, "z_product_detail.html", msg)
}

func Html_MS2018(c *gin.Context) {

	var menu = c.DefaultQuery("n", "product")
	//
	var msg webms.HomePageMenuInfo
	msg.PageTitle = "Web Tool Home Menu"
	msg.RightHtml = fmt.Sprintf("/%s", menu)
	msg.NodeList = webms.GetListAll_MS_Menu()

	c.HTML(http.StatusOK, "home_managementstudio2018.html", msg)

}

func Html_InsertProductExplorer(c *gin.Context) {

	var msg shopapp.PageProductInfo
	//
	var info shopapp.ProductExplorerInfo
	info.ExplorerGUID = "1"
	info.ExplorerName = "标签1"

	shopapp.Insert_ProductExplorer(info)

	c.HTML(http.StatusOK, "z_product_explorerlist.html", msg)

}

func Html_Update_ExplorerTopLevel(c *gin.Context) {

	var txt_explorername = c.PostForm("txt_explorername")
	var hid_explorerguid = c.PostForm("txt_explorerguid")
	//
	shopapp.Update_ProductExplorer_Name(hid_explorerguid, txt_explorername)
	//
	var msg shopapp.PageLabelEditInfo
	msg.PageTitle = "标签编辑"
	msg.LabelList = shopapp.GetListAll_ProductExplorer()
	//
	c.HTML(http.StatusOK, "z_product_explorerlist.html", msg)
}

func Html_ExplorerList_Update(c *gin.Context) {

	var msg shopapp.PageLabelEditInfo
	msg.PageTitle = "标签编辑"

	msg.LabelList = shopapp.GetListAll_ProductExplorer()
	//
	var explorerguid = c.PostForm("txt_explorerguid")
	var lblid = c.PostForm("txt_labelguid")
	var txt_labelname = c.PostForm("txt_labelname")

	for k := range msg.LabelList {
		if msg.LabelList[k].ExplorerGUID == explorerguid {
			//
			for k2 := range msg.LabelList[k].SecondLevelTagList {
				if lblid == msg.LabelList[k].SecondLevelTagList[k2].LabelGUID {
					msg.LabelList[k].SecondLevelTagList[k2].LabelName = txt_labelname
				} else {

				}
			}
			//
			shopapp.Update_ProductExplorer(msg.LabelList[k])
		}
	}
	//
	// msg.LabelList = shopapp.GetListAll_ProductExplorer()

	c.HTML(http.StatusOK, "z_product_explorerlist.html", msg)
}

func GetLabelGUID(lst []shopapp.ProductLabelInfo, explorerGUID string) string {
	dirid, err := strconv.Atoi(explorerGUID)
	if err != nil {
		fmt.Println(err)
	}

	var lblid int = dirid*10000 + 1
	//
	var count = len(lst)
	if count > 0 {

		var a = lst[count-1].LabelGUID
		id, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
		}

		lblid = id + 1
	} else {
		//start
	}
	//
	// fmt.Println(fmt.Sprintf("newid:%d", lblid))
	return fmt.Sprintf("%d", lblid)
}

func GetExplorerGUID(lst []shopapp.ProductExplorerInfo) string {

	var lblid int = 1
	//
	var count = len(lst)
	if count > 0 {

		var a = lst[count-1].ExplorerGUID
		id, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
		}

		lblid = id + 1
	} else {
		//start
	}
	//
	// fmt.Println(fmt.Sprintf("newid:%d", lblid))
	return fmt.Sprintf("%d", lblid)
}

func Html_ExplorerList_AddNew(c *gin.Context) {

	var msg shopapp.PageLabelEditInfo
	msg.PageTitle = "标签编辑"
	msg.LabelList = shopapp.GetListAll_ProductExplorer()
	//
	var dir_1 shopapp.ProductExplorerInfo
	dir_1.ExplorerGUID = GetExplorerGUID(msg.LabelList)
	dir_1.ExplorerName = "新大类"

	shopapp.Insert_ProductExplorer(dir_1)
	//
	msg.LabelList = shopapp.GetListAll_ProductExplorer()
	c.HTML(http.StatusOK, "z_product_explorerlist.html", msg)
}

func Html_ExplorerMenu(c *gin.Context) {

	c.HTML(http.StatusOK, "z_product_explorermenu.html", "")
}

func Html_ExplorerList(c *gin.Context) {
	var cmd = c.DefaultQuery("cmd", "0")
	//
	var msg shopapp.PageLabelEditInfo
	msg.PageTitle = "标签编辑"
	msg.LabelList = shopapp.GetListAll_ProductExplorer()

	if cmd == "0" {
		//show

	}

	if cmd == "1" {
		//add row
		var explorerguid = c.DefaultQuery("explorerguid", "0")
		for k := range msg.LabelList {
			if msg.LabelList[k].ExplorerGUID == explorerguid {
				var newlabelinfo shopapp.ProductLabelInfo
				newlabelinfo.LabelGUID = GetLabelGUID(msg.LabelList[k].SecondLevelTagList, explorerguid)
				newlabelinfo.LabelName = "New Label"
				newlabelinfo.ExplorerGUID = explorerguid
				//
				msg.LabelList[k].SecondLevelTagList = append(msg.LabelList[k].SecondLevelTagList, newlabelinfo)
				//add op
				shopapp.Update_ProductExplorer(msg.LabelList[k])
			}
		}

	}

	if cmd == "2" {
		//delete sub label row
		var explorerguid = c.DefaultQuery("explorerguid", "0")
		var lblid = c.DefaultQuery("id", "0")
		for k := range msg.LabelList {
			if msg.LabelList[k].ExplorerGUID == explorerguid {
				//
				var arr []shopapp.ProductLabelInfo
				for k2 := range msg.LabelList[k].SecondLevelTagList {
					if lblid == msg.LabelList[k].SecondLevelTagList[k2].LabelGUID {

					} else {
						arr = append(arr, msg.LabelList[k].SecondLevelTagList[k2])
					}
				}
				//
				msg.LabelList[k].SecondLevelTagList = arr
				//add op   s.Remove(2)
				shopapp.Update_ProductExplorer(msg.LabelList[k])
			}
		}
		fmt.Println("del:")
		fmt.Println(explorerguid)
		fmt.Println(lblid)
	}

	if cmd == "3" {
		//delete dir row
		var explorerguid = c.DefaultQuery("explorerguid", "0")
		for k := range msg.LabelList {
			if msg.LabelList[k].ExplorerGUID == explorerguid {
				//
				shopapp.Delete_ProductExplorer(msg.LabelList[k])
				msg.LabelList = shopapp.GetListAll_ProductExplorer()
			}
		}
	}

	c.HTML(http.StatusOK, "z_product_explorerlist.html", msg)
}

func Html_ShowProduct_ToPage(c *gin.Context) {

	var gotopagetxt = c.PostForm("txtgotopage")
	gotopage, err := strconv.Atoi(gotopagetxt)
	if err != nil {
		fmt.Println(err)
	}

	var txtmaxpage = c.PostForm("txtmaxpage")
	maxpage, err := strconv.Atoi(txtmaxpage)
	if err != nil {
		fmt.Println(err)
	}

	//
	var msg shopapp.PageProductInfo
	msg.PageTitle = "Product Home"
	msg.MaxPage = maxpage
	msg.GoToPage = gotopage
	msg.PageSize = 6
	msg.CurrentPageIndex = gotopage
	//
	msg.PaginationPre = fmt.Sprintf("/product?op=%s&currentpage=%d&maxpage=%d", "pre", msg.CurrentPageIndex, msg.MaxPage)
	msg.PaginationNext = fmt.Sprintf("/product?op=%s&currentpage=%d&maxpage=%d", "next", msg.CurrentPageIndex, msg.MaxPage)
	//
	msg.CurrentPageIndex = gotopage
	msg.DataList = shopapp.GetList_PagingProduct(msg.CurrentPageIndex, msg.PageSize)
	//
	msg = SetPageStyle(msg)

	c.HTML(http.StatusOK, "z_product_home.html", msg)
}

func Html_ShowProduct(c *gin.Context) {
	//
	var msg shopapp.PageProductInfo
	msg.PageTitle = "Product Home"
	msg.MaxPage = 1
	msg.GoToPage = 1
	msg.PageSize = 6
	msg.CurrentPageIndex = 1
	//
	var op = c.DefaultQuery("op", "default")
	var currentpage = c.DefaultQuery("currentpage", "1")
	curpage, err := strconv.Atoi(currentpage)
	if err != nil {
		fmt.Println(err)
	}

	if curpage == 1 {
		//初始化时，计算一次最大页数
		msg.MaxPage = GetPageCount(shopapp.GetCount_ProductAll(), msg.PageSize)
	} else {
		var maxpagetxt = c.DefaultQuery("maxpage", "1")
		maxpage, err := strconv.Atoi(maxpagetxt)
		if err != nil {
			fmt.Println(err)
		}
		msg.MaxPage = maxpage
	}

	if op == "pre" {
		msg.CurrentPageIndex = curpage - 1
		if msg.CurrentPageIndex <= 0 {
			msg.CurrentPageIndex = 1
		}
		fmt.Println("click pre")
	}

	if op == "next" {
		msg.CurrentPageIndex = curpage + 1
		if msg.CurrentPageIndex > msg.MaxPage {
			msg.CurrentPageIndex = msg.MaxPage
		}
		fmt.Println("click next")
	}

	//
	msg = SetPageStyle(msg)

	//

	msg.PaginationPre = fmt.Sprintf("/product?op=%s&currentpage=%d&maxpage=%d", "pre", msg.CurrentPageIndex, msg.MaxPage)
	msg.PaginationNext = fmt.Sprintf("/product?op=%s&currentpage=%d&maxpage=%d", "next", msg.CurrentPageIndex, msg.MaxPage)
	//
	msg.DataList = shopapp.GetList_PagingProduct(msg.CurrentPageIndex, msg.PageSize)
	//

	c.HTML(http.StatusOK, "z_product_home.html", msg)
}

func SetPageStyle(msg shopapp.PageProductInfo) shopapp.PageProductInfo {
	//
	var styledefault = "a_btn_nextpage_style"
	var styleselect = "a_btn_nextpage_style_select"

	msg.Pagination_1.PageNum = 1
	if msg.CurrentPageIndex == msg.Pagination_1.PageNum {
		msg.Pagination_1.PageNumStyle = styleselect
	} else {
		msg.Pagination_1.PageNumStyle = styledefault
	}

	msg.Pagination_2.PageNum = 2
	if msg.CurrentPageIndex == msg.Pagination_2.PageNum {
		msg.Pagination_2.PageNumStyle = styleselect
	} else {
		msg.Pagination_2.PageNumStyle = styledefault
	}

	msg.Pagination_3.PageNum = 3
	if msg.CurrentPageIndex == msg.Pagination_3.PageNum {
		msg.Pagination_3.PageNumStyle = styleselect
	} else {
		msg.Pagination_3.PageNumStyle = styledefault
	}

	msg.Pagination_4.PageNum = 4
	if msg.CurrentPageIndex == msg.Pagination_4.PageNum {
		msg.Pagination_4.PageNumStyle = styleselect
	} else {
		msg.Pagination_4.PageNumStyle = styledefault
	}

	msg.Pagination_5.PageNum = 5
	if msg.CurrentPageIndex == msg.Pagination_5.PageNum {
		msg.Pagination_5.PageNumStyle = styleselect
	} else {
		msg.Pagination_5.PageNumStyle = styledefault
	}

	msg.Pagination_6.PageNum = 6
	if msg.CurrentPageIndex == msg.Pagination_6.PageNum {
		msg.Pagination_6.PageNumStyle = styleselect
	} else {
		msg.Pagination_6.PageNumStyle = styledefault
	}
	//
	return msg
}

func GetPageCount(rowcount int, pagesize int) int {

	var pagecount = rowcount / pagesize
	var chkrowcount = pagecount * pagesize
	if chkrowcount < rowcount {
		pagecount = pagecount + 1
	}

	return pagecount
}