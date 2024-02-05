package component

type BannerImgComponent struct {
	DataMap    map[string]string
	Source     string
	Parameters map[string]string
	response   []BannerImgResponse
}
type BannerImgResponse struct {
	ImageUrl string `json:"image_url"`
	Ratio    string `json:"ratio"`
	CtaLink  string `json:"cta_link"`
	ImageID  int64  `json:"image_id"`
}

func NewBannerImgComponent(dataMap map[string]string, source string) *BannerImgComponent {
	return &BannerImgComponent{DataMap: dataMap, Source: source}
}

func (bic *BannerImgComponent) SetData(data []map[string]interface{}) {
	for _, v := range data {
		t := BannerImgResponse{}
		for val, key := range bic.DataMap {
			switch key {
			case "image_url":
				t.ImageUrl = v[val].(string)
			case "ratio":
				t.Ratio = "2:1"
			case "cta_link":
				t.CtaLink = v[val].(string)
			}
		}
		bic.response = append(bic.response, t)
	}
}

func (bic *BannerImgComponent) GetData() interface{} {
	return bic.response
}
