package pricemodel

type CreatePrice struct{
    ID      int64   `json:"ID"`
	Price 	int64 	`json:"Price_day"`
}

type UpdatePrice struct{
    ID         int64     `json:"ID"`
    Price      int64     `json:"Price_day"`
}

type GetPrice struct{
    ID int64 `json:"ID"`
}
