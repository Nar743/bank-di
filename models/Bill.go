package models

type Bill struct {
	ID           int    `json:"ID"`
	Number       string `json:"number"`
	Limit        int    `json:"limit"`
	Cards        []Card `json:"cards"`
	IsBillActive bool   `json:"isBillActive"`
}
