package handlers

import (
	"bank-di/models"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomNumber(numb int) string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < numb; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func (h Handler) CreateCard(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var JsonResponce models.Bill
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&JsonResponce)
	if err != nil {
		log.Println(err)
	}
	queryBillId := "SELECT id FROM bills WHERE number = $1"
	bill_id, err := h.DB.Query(queryBillId, JsonResponce.Number)
	if err != nil {
		log.Println(err)
	}

	card := models.Card{
		Number:         randomNumber(16),
		Cvv:            randomNumber(3),
		ExpirationDate: time.Time{},
		Balance:        0,
		History:        nil,
		IsCardActive:   true,
	}
	query := "INSERT INTO cards (bill_id, number, scc, expiration_date, iscardactive, balance) VALUES ($1, $2, $3, $4, $5)"
	_, err = h.DB.Exec(query, bill_id, card.Number, card.Cvv, card.ExpirationDate, card.IsCardActive, card.Balance)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}
