package models

import "time"

type History struct {
	ID                int       `json:"ID"`
	Date              time.Time `json:"date"`
	DestinationCardId string    `json:"destinationCardId"`
	ArrivalCardId     string    `json:"arrivalCardId"`
	OperationType     string    `json:"operationType"`
	Sum               float64   `json:"sum"`
}
