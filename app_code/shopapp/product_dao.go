package shopapp

import (
	"github.com/goago/app_code/db"
	"gopkg.in/mgo.v2/bson"
)

//import "github.com/SharingProduct/sharefund/module/util"

//Layer 1: MongoDB ProductDAO.go
//Database fields are all lowercase,Go code struct field shu

func GetDTName_Product() string {
	return "product"
}

func Insert_Product(info ProductInfo) string {

	c := db.GetMGODB_DT(GetDTName_Product())
	err := c.Insert(&ProductInfo{
		info.ProductGUID,
		info.Name,
		info.IconUrl,
		info.ProductUrl,
		info.ProductDesText,
		info.ProductFrom,
		info.OPTime,
	})
	var suc string
	if err != nil {
		suc = "err"
	} else {
		suc = "suc"
	}
	return suc
}

func Update_Product(info ProductInfo) string {
	c := db.GetMGODB_DT(GetDTName_Product())
	//search conditions
	pkrow := bson.M{"productguid": info.ProductGUID}

	//set value
	newrow := bson.M{"$set": bson.M{
		"productguid":    info.ProductGUID,
		"name":           info.Name,
		"producturl":     info.ProductUrl,
		"producturlfrom": info.ProductFrom,
		"productdestext": info.ProductDesText,
		"optime":         info.OPTime,
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

func Delete_Product(info ProductInfo) {
	c := db.GetMGODB_DT(GetDTName_Product())
	//search conditions
	pkrow := bson.M{"productguid": info.ProductGUID}

	err := c.Remove(pkrow)
	if err != nil {
		//save log
	} else {
	}
}

func GetInfo_Product(productguid string) ProductInfo {
	newrow := ProductInfo{}
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(bson.M{"productguid": productguid}).One(&newrow)
	c.Database.Session.Close()
	return newrow
}

func GetListAll_Product() []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(nil).All(&lst)
	return lst
}

func GetList_Product(productguid string) []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(bson.M{"productguid": productguid}).All(&lst)

	return lst
}

func GetCount_Product() int {
	c := db.GetMGODB_DT(GetDTName_Product())
	n, err := c.Find(nil).Count()
	if err != nil {
	}
	c.Database.Session.Close()
	return n
}
