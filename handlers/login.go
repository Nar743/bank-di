package handlers

import (
	"bank-di/models"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var JsonResponce models.Account
	err = json.Unmarshal(body, &JsonResponce)
	if err != nil {
		log.Fatal(err)
	}
	var account models.Account
	query := "SELECT email, password FROM accounts WHERE email = $1"
	row := h.DB.QueryRow(query, JsonResponce.Email)
	err = row.Scan(&account.Email, &account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if account.Password == JsonResponce.Password {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Entry completed")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Not found Login or password")
	}
}
