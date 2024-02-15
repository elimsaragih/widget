package component

import (
	"encoding/json"
	"log"
)

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

func (bic *BannerImgComponent) SetData(data []byte) error {
	var mapData []map[string]interface{}

	err := json.Unmarshal(data, &mapData)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range mapData {
		t := BannerImgResponse{}
		for val, key := range bic.DataMap {
			switch val {
			case "image_url":
				t.ImageUrl = v[key].(string)
			case "ratio":
				t.Ratio = "2:1"
			case "cta_link":
				t.CtaLink = v[key].(string)
			}
		}
		bic.response = append(bic.response, t)
	}
	return nil
}

func (bic *BannerImgComponent) GetData() interface{} {
	return bic.response
}
