package paymentsmodel

type CreateWork struct {
	Name 		string 	`json:"name_workspace"`
	Date    	string 	`json:"date"`
	Price 		int64   `json:"price"`
	Time 		int32 	`json:"time"`
	Penalty 	int64 	`json:"penalty"`
	UserID  	int64 	`json:"user_id"`
}

type UpdateWork struct {
	ID      	int64   `json:"id"`
	Name 		string 	`json:"name_workspace"`
	Date    	string 	`json:"date"`
	Price 		int64   `json:"price"`
	Time 		int32 	`json:"time"`
	Penalty 	int64 	`json:"penalty"`
}

type DeleteWork struct {
	ID      	int64   `json:"id"`
}

type GetAllWork struct {
    ID      	int64   `json:"id"`
	Name 		string 	`json:"name_workspace"`
	UserID  	int64 	`json:"user_id"`
}

type GetAllWorkByDate struct {
	ID      	int64   `json:"id"`
	Name 		string 	`json:"name_workspace"`
	Date    	string 	`json:"date"`
	UserID  	int64 	`json:"user_id"`
}