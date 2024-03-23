package handler

import "log/slog"



type Payments interface {
	Create()
	Update()
	Get()
	Delete()
}

type Handler struct{
	log *slog.Logger
	p Payments
}

func New(payments Payments, log *slog.Logger) Handler { 
	return Handler{
		log: log,
        p: payments,
	}
}

func (h *Handler) Register(){
	
}