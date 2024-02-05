package source

import "context"

type CampaignSoruce struct {
}

func NewCampaignSource() *CampaignSoruce {
	return &CampaignSoruce{}
}

type CampaignProductList struct {
	ProductID    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int64  `json:"product_price"`
	ProductImage string `json:"product_image"`
	Qty          int    `json:"qty"`
	Link         string `json:"link"`
}

func (as *CampaignSoruce) GetProductList(ctx context.Context, shopID, campaignID int64) (res []CampaignProductList) {
	return []CampaignProductList{
		{
			ProductID:    1,
			ProductName:  "Campaign",
			ProductPrice: 100000,
			ProductImage: "images/campaign",
			Qty:          1,
			Link:         "www.campaign.com",
		},
		{
			ProductID:    2,
			ProductName:  "Campaign2",
			ProductPrice: 100000,
			ProductImage: "images/campaign2",
			Qty:          1,
			Link:         "www.campaign2.com",
		},
	}
}
