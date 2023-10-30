package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (h Handler) TransferFunds(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var transferRequest struct {
		FromAccountID int     `json:"from_account_id"`
		ToAccountID   int     `json:"to_account_id"`
		Amount        float64 `json:"amount"`
	}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&transferRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//проверка на наличие достаточных средств на счете отправителя
	fromAccountQuery := "SELECT balance FROM accounts WHERE id = $1"
	var fromAccountBalance float64
	err = h.DB.QueryRow(fromAccountQuery, transferRequest.FromAccountID).Scan(&fromAccountBalance)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if fromAccountBalance < transferRequest.Amount {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the balances of the from and to accounts
	updateQuery := "UPDATE accounts SET balance = balance - $1 WHERE id = $2"
	_, err = h.DB.Exec(updateQuery, transferRequest.Amount, transferRequest.FromAccountID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = h.DB.Exec(updateQuery, transferRequest.Amount, transferRequest.ToAccountID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
