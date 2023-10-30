package models

import "time"

type Card struct {
	ID             int       `json:"ID" `
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	History        []History `json:"history"`
	IsCardActive   bool      `json:"isCardActive"`
}
