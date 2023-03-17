package pdf

type PageMode Name

const (
	PageModeDontCare       PageMode = ""
	PageModeUseNone        PageMode = "UseNone"
	PageModeUseThumbs      PageMode = "UseThumbs"
	PageModeUseBookmarks   PageMode = "UseBookmarks"
	PageModeFullScreen     PageMode = "FullScreen"
	PageModeUseOC          PageMode = "UseOC"
	PageModeUseAttachments PageMode = "UseAttachments"
)

type PageLayout uint8

const (
	PageLayoutIgnore PageLayout = iota
	PageLayoutDefault
	PageLayoutSinglePage
	PageLayoutOneColumn
	PageLayoutTwoColumnLeft
	PageLayoutTwoColumnRight
	PageLayoutTwoPageLeft
	PageLayoutTwoPageRight
)
