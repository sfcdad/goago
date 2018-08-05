/*
商品类别
*/
package shopapp

type ProductExplorerInfo struct {
	//分类ID
	ExplorerGUID string

	//分类名称
	ExplorerName string

	//状态
	Status int

	//第2级分类 标签列表
	SecondLevelTagList []ProductLabelInfo
}

type ProductLabelInfo struct {

	//1 分类ID
	ExplorerGUID string

	//第2级分类ID
	LabelGUID string

	//第2级分类名称
	LabelName string

	//状态是否选中Checked
	Status int

	//html css name
	ClassStyle string

	//产品GUID
	ProductGUID string
}
