package paymentsmodel

type ListPayments struct {
	ID       int    `json:"id"`
	NameWorkspace string `json:"name_workspace"`
	Date     string `json:"date"`
	PriceDay int    `json:"price"`
	Token    string `json:"token"`
}
