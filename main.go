package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	components "github.com/elimsaragih/widget/component"
	source "github.com/elimsaragih/widget/source"
)

type SourceConfig struct {
	Widget    string            `json:"widget"`
	Component string            `json:"component"`
	Source    string            `json:"source"`
	Param     string            `json:"param"`
	Mapping   map[string]string `json:"mapping"`
}

func main() {

	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	var appConfig []SourceConfig

	if err := json.Unmarshal(configData, &appConfig); err != nil {
		fmt.Println(string(configData))
		fmt.Println("Error parsing configuration:", err)
		return
	}

	var widgets []Widget

	// generate widget base on component
	for i, v := range appConfig {
		switch v.Component {
		case "banner":
			banner := components.NewBannerImgComponent(v.Mapping, v.Source)
			temp := Widget{
				Header: Header{
					Title: "title: " + v.Widget,
				},
				Body: Body{
					Components: []Component{
						{
							Identifier: v.Widget,
							Source:     v.Source,
							Styles: []Style{
								{
									Key:   "position",
									Value: strconv.Itoa(i + 1),
								},
							},
							Data: banner,
						},
					},
				},
			}
			widgets = append(widgets, temp)
		case "product_list":
			banner := components.NewProductListComponent(v.Mapping, v.Source)
			temp := Widget{
				Header: Header{
					Title: "title: " + v.Widget,
				},
				Body: Body{
					Components: []Component{
						{
							Identifier: v.Widget,
							Source:     v.Source,
							Styles: []Style{
								{
									Key:   "position",
									Value: strconv.Itoa(i + 1),
								},
							},
							Data: banner,
						},
					},
				},
			}
			widgets = append(widgets, temp)
		}
	}

	// contract widget data by source
	for _, v := range widgets {
		for _, comp := range v.Body.Components {
			switch comp.Source {
			case "ace":
				temp := source.NewAceSource()
				tes := AdaptorAce(temp.GetProductList(context.Background(), 123))
				comp.Data.SetData(tes)
			case "campaign":
				temp := source.NewCampaignSource()
				tes := AdaptorCampaign(temp.GetProductList(context.Background(), 123, 123))
				comp.Data.SetData(tes)
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
