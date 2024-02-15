package widgetmaster

// Creation Pattern Factory

// WidgetMaster factory
type WidgetMaster struct {
	Header Header
	Body   Body
}

func InitWidget(setupData ComponentData, source SourceConfig) WidgetMaster {
	return WidgetMaster{
		Header: Header{
			Title: "title: " + source.Widget,
		},
		Body: Body{
			Components: []Component{
				{
					Identifier: source.Widget,
					Source:     source.Source,
					Styles:     []Style{},
					Data:       setupData,
				},
			},
		},
	}
}

type Header struct {
	Title    string
	SubTitle string
	CtaText  string
	CtaLnk   string
}

type Body struct {
	Components []Component
}

type Style struct {
	Key   string
	Value string
}

type Component struct {
	Identifier string
	Source     string
	Styles     []Style
	Data       ComponentData
}

type ComponentData interface {
	SetData(data []byte) error
	GetData() interface{}
}
