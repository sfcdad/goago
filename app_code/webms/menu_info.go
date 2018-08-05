package webms

type MS_MenuInfo struct {
	MenuGUID string
	MenuName string
	NodeList []MenuInfoNode
}

type MenuInfoNode struct {
	MenuGUID string
	MenuName string
	MenuUrl  string
}
