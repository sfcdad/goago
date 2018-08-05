package shopapp

import (
	"fmt"
	"github.com/goago/app_code/db"
	"github.com/goago/app_code/common"
	"gopkg.in/mgo.v2/bson"
)

//import "github.com/SharingProduct/sharefund/module/util"

//Layer 1: MongoDB ProductDAO.go
//Database fields are all lowercase,Go code struct field shu

func GetDTName_Product() string {
	return "product"
}

func Insert_ProductEmptyRow(info ProductInfo) string {
	var suc string
	suc = Insert_Product(info)
	return suc
}

func Insert_ProductFromAPI(info ProductInfo) string {
	var suc string
	//
	var count = GetCount_Product(info.SKUID)
	if count > 0 {
		suc = "exist"
		return suc
	}
	//
	suc = Insert_Product(info)
	return suc
}

func Insert_Product(info ProductInfo) string {
	var suc string
	//
	c := db.GetMGODB_DT(GetDTName_Product())
	err := c.Insert(&ProductInfo{
		info.ExplorerGUID,
		info.ExplorerName,
		info.ProductGUID,
		info.ProductName,
		info.IconUrl,
		info.ProductUrl,
		info.ProductDesText,
		info.ProductComment,
		info.SKUID,
		info.ProductFrom,
		info.OPTime,
		info.OPTimeText,
		info.SellUrl,
		info.StatusID,
		info.StatusText,
	})

	if err != nil {
		suc = "err"
	} else {
		suc = "suc"
	}
	return suc
}

func Update_ProductSearchLabel(productguid string, lbls string, names string) string {
	c := db.GetMGODB_DT(GetDTName_Product())
	//search conditions
	pkrow := bson.M{"productguid": productguid}
	//set value
	newrow := bson.M{"$set": bson.M{
		"explorerguid": lbls,
		"explorername": names,
		"optimetext":   common.GetBeiJingTime(),
		"optime":       common.GetNowTimestamp(),
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

func Update_Product(info ProductInfo) string {
	c := db.GetMGODB_DT(GetDTName_Product())
	//search conditions
	pkrow := bson.M{"productguid": info.ProductGUID}

	//set value
	newrow := bson.M{"$set": bson.M{
		"productname":    info.ProductName,
		"iconurl":        info.IconUrl,
		"sellurl":        info.SellUrl,
		"productdestext": info.ProductDesText,
		"optimetext":     info.OPTimeText,
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

func Update_Product_ProductComment(ProductGUID string,ProductComment string) string {
	c := db.GetMGODB_DT(GetDTName_Product())
	//
	var info ProductInfo
	info.ProductGUID = ProductGUID
	info.ProductComment = ProductComment
	//search conditions
	pkrow := bson.M{"productguid": info.ProductGUID}

	//set value
	info.OPTime = common.GetNowTimestamp()
	info.OPTimeText = common.GetBeiJingTime()
	newrow := bson.M{"$set": bson.M{
		"productcomment":    info.ProductComment,
		"optimetext":     info.OPTimeText,
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

func Update_Product_ProductDesText(ProductGUID string,ProductDesText string) string {
	c := db.GetMGODB_DT(GetDTName_Product())
	//
	var info ProductInfo
	info.ProductGUID = ProductGUID
	info.ProductDesText = ProductDesText
	//search conditions
	pkrow := bson.M{"productguid": info.ProductGUID}

	//set value
	info.OPTime = common.GetNowTimestamp()
	info.OPTimeText = common.GetBeiJingTime()
	newrow := bson.M{"$set": bson.M{
		"productdestext":    info.ProductDesText,
		"optimetext":     info.OPTimeText,
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

func GetCount_Product(SKUID int) int {
	c := db.GetMGODB_DT(GetDTName_Product())
	//
	n, err := c.Find(bson.M{"skuid": SKUID}).Count()
	if err != nil {
	}
	c.Database.Session.Close()
	return n
}

func GetCount_ProductAll() int {
	c := db.GetMGODB_DT(GetDTName_Product())
	//
	n, err := c.Find(nil).Count()
	if err != nil {
	}
	c.Database.Session.Close()
	return n
}

func GetInfo_Product(productguid string) ProductInfo {
	newrow := ProductInfo{}
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(bson.M{"productguid": productguid}).One(&newrow)
	c.Database.Session.Close()
	return newrow
}
func GetList_Product_ByLabel(lbls string) []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	//模糊查询，比如参数是：11 记录里面11,  11,13
	c.Find(bson.M{"explorerguid": bson.M{"$regex": lbls}}).All(&lst)
	return lst
}

func GetListAll_Product() []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(nil).All(&lst)
	return lst
}

func GetList_PagingProduct(curpage int, pagesize int) []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	// c.Find(nil).All(&lst)
	// 表示从偏移位置为2的地方开始取两条记录，索引2就是第3条记录
	var rowindex = 0
	if curpage > 0 {
		rowindex = (curpage - 1) * pagesize
	}
	err := c.Find(nil).Sort("-optime").Skip(rowindex).Limit(pagesize).All(&lst) //-optime 降序
	if err != nil {
		fmt.Println(err)
	}
	return lst
}

func GetList_Product(productguid string) []ProductInfo {
	var lst []ProductInfo
	c := db.GetMGODB_DT(GetDTName_Product())
	c.Find(bson.M{"productguid": productguid}).All(&lst)

	return lst
}

// func GetCount_Product() int {
// 	c := db.GetMGODB_DT(GetDTName_Product())
// 	n, err := c.Find(nil).Count()
// 	if err != nil {
// 	}
// 	c.Database.Session.Close()
// 	return n
// }
