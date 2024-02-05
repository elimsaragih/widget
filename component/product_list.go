package component

type ProductListComponent struct {
	DataMap    map[string]string
	Source     string
	Parameters map[string]string
	response   []ProductListResponse
}
type ProductListResponse struct {
	ProductID int64  `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	CtaLink   string `json:"cta_link"`
	Price     int64  `json:"price"`
}

func NewProductListComponent(dataMap map[string]string, source string) *ProductListComponent {
	return &ProductListComponent{DataMap: dataMap, Source: source}
}

func (bic *ProductListComponent) SetData(data []map[string]interface{}) {
	for _, v := range data {
		t := ProductListResponse{}
		for val, key := range bic.DataMap {
			switch val {
			case "product_id":
				t.ProductID = v[key].(int64)
			case "image_url":
				t.ImageUrl = v[key].(string)
			case "cta_link":
				t.CtaLink = v[key].(string)
			case "price":
				t.Price = v[key].(int64)
			}
		}
		bic.response = append(bic.response, t)
	}
}

func (bic *ProductListComponent) GetData() interface{} {
	return bic.response
}
