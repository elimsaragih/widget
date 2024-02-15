package widgetmaster

// Creation Pattern Factory

// WidgetMaster factory
type WidgetMaster struct {
	Header Header
	Body   Body
}

func InitWidget(setupData ComponentData, title, source string) *WidgetMaster {
	return &WidgetMaster{
		Header: Header{
			Title: "title: " + title,
		},
		Body: Body{
			Components: []Component{
				{
					Source: source,
					Styles: []Style{},
					Data:   setupData,
				},
			},
		},
	}
}

func (w *WidgetMaster) SetHeader(header Header) {
	w.Header = header
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
