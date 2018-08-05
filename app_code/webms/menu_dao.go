package webms

import (
	"github.com/goago/app_code/db"
	"gopkg.in/mgo.v2/bson"
)

//Layer 1: MongoDB MS_MenuDAO.go
//Database fields are all lowercase,Go code struct field shu
func GetDTName_MS_Menu() string {
	return "ms_menu"
}

func InsertTestData_Menu(){

	var noderoot_1 MS_MenuInfo
	noderoot_1.MenuGUID = "1"
	noderoot_1.MenuName = "商品管理"

	var row1 MenuInfoNode
	row1.MenuGUID = "11"
	row1.MenuName = "商品列表"
	row1.MenuUrl = "/ms?n=product"

	var row2 MenuInfoNode
	row2.MenuGUID = "12"
	row2.MenuName = "新增商品"
	row2.MenuUrl = "/ms?n=product_add"

	var row3 MenuInfoNode
	row3.MenuGUID = "13"
	row3.MenuName = "商品分类"
	row3.MenuUrl = "/ms?n=explorermenu"

	noderoot_1.NodeList = append(noderoot_1.NodeList, row1)
	noderoot_1.NodeList = append(noderoot_1.NodeList, row2)
	noderoot_1.NodeList = append(noderoot_1.NodeList, row3)


	Insert_MS_Menu(noderoot_1)
}

func Insert_MS_Menu(info MS_MenuInfo) string {

	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	err := c.Insert(&MS_MenuInfo{
		info.MenuGUID,
		info.MenuName,
		info.NodeList,
	})
	var suc string
	if err != nil {
		suc = "err"
	} else {
		suc = "suc"
	}
	return suc
}

func Update_MS_Menu(info MS_MenuInfo) string {
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	//search conditions
	pkrow := bson.M{"menuguid": info.MenuGUID}

	//set value
	newrow := bson.M{"$set": bson.M{
		"menuguid": info.MenuGUID,
		"menuname": info.MenuName,
		"nodelist": info.NodeList,
	}}

	err := c.Update(pkrow, newrow)
	var suc string
	if err != nil {
		suc = "err"
	} else {
		suc = "suc"
	}
	return suc
}

func Delete_MS_Menu(info MS_MenuInfo) {
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	//search conditions
	pkrow := bson.M{"menuguid": info.MenuGUID}

	err := c.Remove(pkrow)
	if err != nil {
		//save log
	} else {
	}
}

func GetInfo_MS_Menu(menuguid string) MS_MenuInfo {
	newrow := MS_MenuInfo{}
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	c.Find(bson.M{"menuguid": menuguid}).One(&newrow)
	c.Database.Session.Close()
	return newrow
}

func GetListAll_MS_Menu() []MS_MenuInfo {
	var lst []MS_MenuInfo
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	c.Find(nil).All(&lst)
	return lst
}

func GetList_MS_Menu(menuguid string) []MS_MenuInfo {
	var lst []MS_MenuInfo
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	c.Find(bson.M{"menuguid": menuguid}).All(&lst)

	return lst
}

func GetCount_MS_Menu() int {
	c := db.GetMGODB_DT(GetDTName_MS_Menu())
	n, err := c.Find(nil).Count()
	if err != nil {
	}
	c.Database.Session.Close()
	return n
}
