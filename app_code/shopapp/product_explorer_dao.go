package shopapp

import (
	"github.com/goago/app_code/db"
	"gopkg.in/mgo.v2/bson"
)

//Layer 1: MongoDB ProductExplorerDAO.go
//Database fields are all lowercase,Go code struct field shu

func GetDTName_ProductExplorer() string {
	return "productexplorer"
}

func InitTestData_ProductExplorer() {
	var dir_1 ProductExplorerInfo
	dir_1.ExplorerGUID = "1"
	dir_1.ExplorerName = "母婴"

	var lbl_1 ProductLabelInfo
	lbl_1.LabelGUID = "11"
	lbl_1.LabelName = "奶粉"
	lbl_1.ClassStyle = "a_btn_labeldefault"
	dir_1.SecondLevelTagList = append(dir_1.SecondLevelTagList, lbl_1)

	var lbl_2 ProductLabelInfo
	lbl_2.LabelGUID = "12"
	lbl_2.LabelName = "纸尿裤"
	lbl_2.ClassStyle = "a_btn_labeldefault"
	dir_1.SecondLevelTagList = append(dir_1.SecondLevelTagList, lbl_2)

	var lbl_3 ProductLabelInfo
	lbl_3.LabelGUID = "13"
	lbl_3.LabelName = "辅食"
	lbl_3.ClassStyle = "a_btn_labeldefault"
	dir_1.SecondLevelTagList = append(dir_1.SecondLevelTagList, lbl_3)

	//
	Insert_ProductExplorer(dir_1)

	var dir_2 ProductExplorerInfo
	dir_2.ExplorerGUID = "2"
	dir_2.ExplorerName = "美妆护肤"

	var lbl_21 ProductLabelInfo
	lbl_21.LabelGUID = "21"
	lbl_21.LabelName = "护肤"
	lbl_21.ClassStyle = "a_btn_labeldefault"
	dir_2.SecondLevelTagList = append(dir_2.SecondLevelTagList, lbl_21)

	var lbl_22 ProductLabelInfo
	lbl_22.LabelGUID = "22"
	lbl_22.LabelName = "美体"
	lbl_22.ClassStyle = "a_btn_labeldefault"
	dir_2.SecondLevelTagList = append(dir_2.SecondLevelTagList, lbl_22)

	var lbl_23 ProductLabelInfo
	lbl_23.LabelGUID = "23"
	lbl_23.LabelName = "美发"
	lbl_23.ClassStyle = "a_btn_labeldefault"
	dir_2.SecondLevelTagList = append(dir_2.SecondLevelTagList, lbl_23)

	//
	Insert_ProductExplorer(dir_2)

}

func Insert_ProductExplorer(info ProductExplorerInfo) string {

	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	err := c.Insert(&ProductExplorerInfo{
		info.ExplorerGUID,
		info.ExplorerName,
		info.Status,
		info.SecondLevelTagList,
	})
	var suc string
	if err != nil {
		suc = "err"
	} else {
		suc = "suc"
	}
	return suc
}

func Update_ProductExplorer_Name(explorerguid string, name string) string {
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	//search conditions
	pkrow := bson.M{"explorerguid": explorerguid}

	var info = GetInfo_ProductExplorer(explorerguid)
	info.ExplorerName = name
	//set value
	newrow := bson.M{"$set": bson.M{
		"explorername": info.ExplorerName,
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

func Update_ProductExplorer(info ProductExplorerInfo) string {
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	//search conditions
	pkrow := bson.M{"explorerguid": info.ExplorerGUID}

	//set value
	newrow := bson.M{"$set": bson.M{
		"explorername":       info.ExplorerName,
		"status":             info.Status,
		"secondleveltaglist": info.SecondLevelTagList,
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

func Delete_ProductExplorer(info ProductExplorerInfo) {
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	//search conditions
	pkrow := bson.M{"explorerguid": info.ExplorerGUID}

	err := c.Remove(pkrow)
	if err != nil {
		//save log
	} else {
	}
}

func GetInfo_ProductExplorer(explorerguid string) ProductExplorerInfo {
	newrow := ProductExplorerInfo{}
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	c.Find(bson.M{"explorerguid": explorerguid}).One(&newrow)
	c.Database.Session.Close()
	return newrow
}

func GetListAll_ProductExplorer() []ProductExplorerInfo {
	var lst []ProductExplorerInfo
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	c.Find(nil).All(&lst)
	return lst
}

func GetList_ProductExplorer(explorerguid string) []ProductExplorerInfo {
	var lst []ProductExplorerInfo
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	c.Find(bson.M{"explorerguid": explorerguid}).All(&lst)

	return lst
}

func GetCount_ProductExplorer() int {
	c := db.GetMGODB_DT(GetDTName_ProductExplorer())
	n, err := c.Find(nil).Count()
	if err != nil {
	}
	c.Database.Session.Close()
	return n
}
