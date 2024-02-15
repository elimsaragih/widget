package route

import (
	"encoding/json"

	masterWidget "github.com/elimsaragih/widget/widget-master"
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	route  *httprouter.Router
	widget *masterWidget.WidgetMaster
}

func InitShopPagePkg(route *httprouter.Router, widget masterWidget.ComponentData, title, path string, call httprouter.Handle) *Route {
	// banner := components.NewBannerImgComponent(map[string]string{}, "external")
	// widget := masterWidget.InitWidget(banner, masterWidget.AppConfig[1])
	route.GET("/external/"+path, call)
	return &Route{route: route, widget: masterWidget.InitWidget(widget, title, path)}
}

func (r *Route) SetData(data []byte) error {
	for _, comp := range r.widget.Body.Components {
		err := comp.Data.SetData(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Route) GetData() ([]byte, error) {
	return json.Marshal(r.widget)
}

func (r *Route) SetHeaderWidget(header masterWidget.Header) {
	r.widget.Header = header
}
