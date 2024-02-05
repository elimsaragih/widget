package source

import "context"

type AceSource struct {
}

func NewAceSource() *AceSource {
	return &AceSource{}
}

type AceProductList struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

func (as *AceSource) GetProductList(ctx context.Context, shopID int64) (res []AceProductList) {
	return []AceProductList{
		{
			ID:    1,
			Name:  "test",
			Price: 10000,
			Image: "www.test.com",
			Link:  "http://www.test.com",
		},
		{
			ID:    2,
			Name:  "test2",
			Price: 10000,
			Image: "www.test2.com",
			Link:  "http://www.test2.com",
		},
	}
}
