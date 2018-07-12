package shopapp

type HomePageInfo struct {
	PageTitle string
	DataList []ProductInfo
}

func GetInfo_HomePage() HomePageInfo {

	var info HomePageInfo
	info.PageTitle="GoaGo"
	info.DataList = GetListAll_Product()
	//
	return info
}