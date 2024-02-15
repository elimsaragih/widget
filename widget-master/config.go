package widgetmaster

type SourceConfig struct {
	Widget    string            `json:"widget"`
	Component string            `json:"component"`
	Source    string            `json:"source"`
	Param     string            `json:"param"`
	Mapping   map[string]string `json:"mapping"`
}

var AppConfig []SourceConfig
