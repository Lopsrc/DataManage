package listprices

type CreateListPrices struct{
	Date     string `json:"Date"`
	Price 	 int 	`json:"Price_day"`
}