package dto

type KLineParam struct {
	CompareList []string `json:"compare_list"`
	OptionName string `json:"option_name"`
	IsBlock	string `json:"is_block"`
}

type KLineResult struct {
	Colors []string `json:"colors"`
	Data []KLineItem `json:"data"`
	Items []KLineStyle `json:"items"`
}

type KLineItem struct {
	Day string `json:"day"`
	Price float64 `json:"price"`
	Symbol string `json:"symbol"`
}

type KLineStyle struct {
	Name string `json:"name"`
	Marker KLineMarker `json:"marker"`
}

type KLineMarker struct {
	Symbol string `json:"symbol"`
	Style KLineColor `json:"style"`
}

type KLineColor struct {
	Fill string `json:"fill"`
}