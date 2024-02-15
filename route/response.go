package route

type HttpResponse struct {
	Status
	WidgetExtMaster
}

type Status struct {
	Code    string
	Message string
}
type WidgetExtMaster struct {
	Header Header
	Body   Body
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
	Data       string
}
