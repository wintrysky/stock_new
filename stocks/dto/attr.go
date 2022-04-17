package dto

type AttributeParam struct {
	IDs []string `json:"ids"`
	BuyTags string `json:"buy_tags"`
	CompanyTags string `json:"company_tags"`
	Description string `json:"description"`
	ProfitTags string `json:"profit_tags"`
}
