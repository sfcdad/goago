/*
商品详情
*/
package shopapp

type ProductInfo struct {

	//1 分类ID,打上标签,可以多选
	ExplorerGUID string `json:"explorerguid"`

	//2 分类名称
	ExplorerName string `json:"explorername"`

	//3 产品GUID
	ProductGUID string `json:"productguid"`

	//4 商品名称
	ProductName string `json:"productname"`

	//5 展示图标地址
	IconUrl string `json:"iconurl"`

	//6 jd website 商品url-jd
	ProductUrl string `json:"producturl"`

	//7 产品描述详细
	ProductDesText string `json:"productdestext"`

	//8 产品评论
	ProductComment string `json:"productcomment"`

	//9 电商商城商品编号
	SKUID int `json:"skuid"`

	//10 商品url来源-jd,taobaoke,other
	ProductFrom int `json:"productfrom"`

	//11 更新时间
	OPTime int64 `json:"optime"`

	//12 更新时间Text
	OPTimeText string `json:"optimetext"`

	//13 推广代码
	SellUrl string `json:"sellurl"`

	//14
	StatusID int `json:"statusid"`

	//15
	StatusText string `json:"statustext"`
}

type PageLabelEditInfo struct {
	PageTitle        string
	RowInfo          ProductInfo
	LabelList        []ProductExplorerInfo
	CurrentLabelList string
}

type PageProductEditInfo struct {
	PageTitle        string
	RowInfo          ProductInfo
	LabelList        []ProductExplorerInfo
	CurrentLabelList string
}

type PageProductInfo struct {
	PageTitle string        `json:"pagetitle"`
	DataList  []ProductInfo `json:"datalist"`
	RowInfo   ProductInfo   `json:"rowinfo"`

	//上一页
	PaginationPre string

	//下一页
	PaginationNext string

	//到第几页
	PaginationGo string

	//到第几页确定
	PaginationGoOK string

	//页1
	Pagination_1 PageNumInfo

	//页2
	Pagination_2 PageNumInfo

	//页3
	Pagination_3 PageNumInfo

	//页4
	Pagination_4 PageNumInfo

	//页5
	Pagination_5 PageNumInfo

	//页6
	Pagination_6 PageNumInfo

	//当前页
	CurrentPageIndex int

	//最大页数
	MaxPage int

	//跳转到页
	GoToPage int

	//1 page rows count
	PageSize int
}

type PageNumInfo struct {
	PageNum      int
	PageNumStyle string
}
