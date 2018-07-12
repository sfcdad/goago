/*
商品详情
*/
package shopapp

type ProductInfo struct {
	//1 京东产品GUID
	ProductGUID string

	//商品名称
	Name string

	//展示图标地址
	IconUrl string

	//jd website 商品url-jd
	ProductUrl string

	//产品描述详细
	ProductDesText string

	//商品url来源-jd,taobaoke,other
	ProductFrom int

	//更新时间
	OPTime int64
}

type PageProductInfo struct {
	PageTitle string
	DataList  []ProductInfo
}
