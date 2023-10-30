package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (h Handler) SetCardLimit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var limitRequest struct {
		CardID int `json:"card_id"`
		Limit  int `json:"limit"`
	}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&limitRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := "UPDATE cards SET balance = $1 WHERE id = $2"
	_, err = h.DB.Exec(query, limitRequest.Limit, limitRequest.CardID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
