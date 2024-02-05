package main

import (
	"reflect"

	source "github.com/elimsaragih/widget/source"
)

func jsonCasting(refVal reflect.Value) map[string]interface{} {
	t := make(map[string]interface{}, 0)
	for i := 0; i < refVal.NumField(); i++ {
		refType := refVal.Type()
		key := refType.Field(i).Tag.Get("json")
		if key == "" {
			continue
		}

		fieldVal := refVal.Field(i)
		// do not process nil value
		if fieldVal.Kind() == reflect.Ptr {
			if refVal.Field(i).IsNil() {
				continue
			} else {
				// get the non pointer value
				fieldVal = fieldVal.Elem()
			}
		}

		t[key] = fieldVal.Interface()
	}
	return t
}

func AdaptorAce(temp []source.AceProductList) (res []map[string]interface{}) {
	for _, v := range temp {
		refVal := reflect.ValueOf(v)

		res = append(res, jsonCasting(refVal))
	}
	return
}

func AdaptorCampaign(temp []source.CampaignProductList) (res []map[string]interface{}) {
	for _, v := range temp {
		refVal := reflect.ValueOf(v)

		res = append(res, jsonCasting(refVal))
	}
	return
}
