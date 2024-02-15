package route

import (
	"encoding/json"
	"net/http"

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
	obj := &Route{route: route, widget: masterWidget.InitWidget(widget, title, path)}
	route.GET("/external/"+path, wrapCall(obj, call))
	return obj
}

func wrapCall(rObj *Route, call httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		call(w, r, ps)
		jData, _ := json.Marshal(rObj.generateResponse())
		w.Write(jData)
	}
}
func (r *Route) generateResponse() (resp HttpResponse) {
	var components []Component
	for _, v := range r.widget.Body.Components {
		data, _ := json.Marshal(v.Data.GetData())
		temp := Component{
			Identifier: v.Identifier,
			Source:     v.Source,
			Data:       string(data),
		}

		for _, st := range v.Styles {
			temp.Styles = append(temp.Styles, Style(st))
		}
		components = append(components, temp)
	}

	resp.WidgetExtMaster = WidgetExtMaster{
		Header: Header(r.widget.Header),
		Body: Body{
			Components: components,
		},
	}
	return resp
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

func (r *Route) SetHeaderWidget(header masterWidget.Header) {
	r.widget.Header = header
}
