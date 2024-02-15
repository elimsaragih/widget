package component

import (
	"encoding/json"
	"log"
)

type ProductListComponent struct {
	DataMap     map[string]string
	Source      string
	Parameters  map[string]string
	dataMapType int
	response    []ProductListResponse
}
type ProductListResponse struct {
	ProductID int64  `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	CtaLink   string `json:"cta_link"`
	Price     int64  `json:"price"`
}

func NewProductListComponent(dataMapType int, dataMap map[string]string, source string) *ProductListComponent {
	return &ProductListComponent{dataMapType: dataMapType, DataMap: dataMap, Source: source}
}

func (bic *ProductListComponent) SetData(data []byte) error {
	if bic.dataMapType == DataMapDefault {
		return json.Unmarshal(data, &bic.response)
	}

	var mapData []map[string]interface{}
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range mapData {
		t := ProductListResponse{}
		for val, key := range bic.DataMap {
			switch val {
			case "product_id":
				t.ProductID = int64(v[key].(float64))
			case "image_url":
				t.ImageUrl = v[key].(string)
			case "cta_link":
				t.CtaLink = v[key].(string)
			case "price":
				t.Price = int64(v[key].(float64))
			}
		}
		bic.response = append(bic.response, t)
	}
	return nil
}

func (bic *ProductListComponent) GetData() interface{} {
	return bic.response
}
