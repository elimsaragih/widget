package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	components "github.com/elimsaragih/widget/component"
	exterWidget "github.com/elimsaragih/widget/route"
	source "github.com/elimsaragih/widget/source"
	masterWidget "github.com/elimsaragih/widget/widget-master"
)

type RegExternalWidget struct {
	Source string
	Path   string
}

var extList = []string{"campaign"}
var baseURL = "http://localhost:8081"

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

	var widgets []*masterWidget.WidgetMaster
	var widgetsExt []exterWidget.WidgetExtMaster

	// generate widget base on component
	for _, v := range masterWidget.AppConfig {
		switch v.Component {
		case "banner":
			banner := components.NewBannerImgComponent(components.DataMapConfigurable, v.Mapping, v.Source)
			widgets = append(widgets, masterWidget.InitWidget(banner, v.Title, v.Source))
		case "product_list":
			list := components.NewProductListComponent(components.DataMapConfigurable, v.Mapping, v.Source)
			widgets = append(widgets, masterWidget.InitWidget(list, v.Title, v.Source))
		}
	}

	// any vaidation widgets here

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

	for _, v := range extList {
		temp, err := fetchWidget(v)
		if err != nil {
			log.Printf("failed to fetch external widget: %v", err)
			continue
		}
		widgetsExt = append(widgetsExt, temp)
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

	// widget external result
	fmt.Println("External Widget result")
	for _, v := range widgetsExt {
		fmt.Printf("Header: %s", v.Header)
		for _, comp := range v.Body.Components {
			fmt.Printf("Style: %+v", comp.Styles)
			fmt.Printf("Data: %+v", comp.Data)
		}
		fmt.Println("")
	}

}

func fetchWidget(path string) (exterWidget.WidgetExtMaster, error) {
	var err error
	var client = &http.Client{}
	var data exterWidget.WidgetExtMaster

	request, err := http.NewRequest("GET", baseURL+"/external/"+path, nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
