package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	components "github.com/elimsaragih/widget/component"
	source "github.com/elimsaragih/widget/source"
	masterWidget "github.com/elimsaragih/widget/widget-master"
)

func main() {

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	if err := json.Unmarshal(configData, &masterWidget.AppConfig); err != nil {
		fmt.Println(string(configData))
		fmt.Println("Error parsing configuration:", err)
		return
	}

	var widgets []masterWidget.WidgetMaster

	// generate widget base on component
	for _, v := range masterWidget.AppConfig {
		switch v.Component {
		case "banner":
			banner := components.NewBannerImgComponent(v.Mapping, v.Source)
			widgets = append(widgets, masterWidget.InitWidget(banner, v))
		case "product_list":
			list := components.NewProductListComponent(v.Mapping, v.Source)
			widgets = append(widgets, masterWidget.InitWidget(list, v))
		}
	}

	// contract widget data by source
	for _, v := range widgets {
		for _, comp := range v.Body.Components {
			switch comp.Source {
			case "ace":
				temp := source.NewAceSource()
				tes := AdaptorAce(temp.GetProductList(context.Background(), 123))
				data, _ := json.Marshal(tes)
				comp.Data.SetData(data)
			case "campaign":
				temp := source.NewCampaignSource()
				tes := AdaptorCampaign(temp.GetProductList(context.Background(), 123, 123))
				data, _ := json.Marshal(tes)
				comp.Data.SetData(data)
			}
		}
	}

	// widget result
	for _, v := range widgets {
		fmt.Printf("Header: %s", v.Header)
		for _, comp := range v.Body.Components {
			fmt.Printf("Style: %+v", comp.Styles)
			fmt.Printf("Data: %+v", comp.Data.GetData())
		}
		fmt.Println("")
	}

}
