package paymentsmodel

import "github.com/jackc/pgtype"

type Work struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name_workspace"`
	Date    pgtype.Date `json:"date"`
	Price   int64       `json:"price"`
	Time    int32       `json:"time"`
	Penalty int64       `json:"penalty"`
	UserID  int64       `json:"user_id"`
}
