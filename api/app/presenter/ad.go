package presenter

import "hbbapi/app/helper"

type AdBase struct {
	Id     int    `json:"id"`
	AdLink string `json:"ad_link"`
	AdCode string `json:"ad_code"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func FormatAdBase(data interface{}) []AdBase {
	var adBase []AdBase
	helper.StructToStruct(data, &adBase)
	return adBase
}
