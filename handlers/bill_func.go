package handlers

import (
	"bank-di/models"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomNumberBill() string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func (h Handler) CreateBill(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users models.Account
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := "SELECT * FROM accounts WHERE email = $1"
	row := h.DB.QueryRow(query, users.Email)

	var dbUser models.Account
	err = row.Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.SecondName, &dbUser.Email, &dbUser.Password)
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

	if users.Password != dbUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var bill = models.Bill{
		ID:           dbUser.ID,
		Number:       randomNumberBill(),
		Limit:        0,
		Cards:        nil,
		IsBillActive: true,
	}

	query = "INSERT INTO bills (account_id, number, sum_limit, is_bill_active) VALUES ($1, $2, $3, $4)"
	_, err = h.DB.Exec(query, bill.ID, bill.Number, bill.Limit, bill.IsBillActive)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bill)
}

func (h Handler) CloseBill(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data struct {
		AccountID int    `json:"account_id"`
		Password  string `json:"password"`
		BillID    int    `json:"bill_id"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := "SELECT * FROM accounts WHERE id = $1"
	row := h.DB.QueryRow(query, data.AccountID)

	var dbUser models.Account
	err = row.Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.SecondName, &dbUser.Email, &dbUser.Password)
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

	if data.Password != dbUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	query = "UPDATE bills SET is_bill_active = false WHERE id = $1 AND account_id = $2"
	_, err = h.DB.Exec(query, data.BillID, data.AccountID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetBillInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data struct {
		AccountID int    `json:"account_id"`
		Password  string `json:"password"`
		BillID    int    `json:"bill_id"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := "SELECT * FROM accounts WHERE id = $1"
	row := h.DB.QueryRow(query, data.AccountID)

	var dbUser models.Account
	err = row.Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.SecondName, &dbUser.Email, &dbUser.Password)
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

	if data.Password != dbUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	query = "SELECT * FROM bills WHERE id = $1 AND account_id = $2"
	row = h.DB.QueryRow(query, data.BillID, data.AccountID)

	var bill models.Bill
	err = row.Scan(&bill.ID, &bill.Number, &bill.Limit, &bill.IsBillActive)
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bill)
}

//func randomNumberBill() string {
//	var number string
//	source := rand.NewSource(time.Now().UnixNano())
//	r := rand.New(source)
//	for i := 0; i < 20; i++ {
//		number += strconv.Itoa(r.Intn(10))
//	}
//	return number
//}
//
//func (h Handler) CreateBill(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Println(err)
//	}
//
//	var users models.Account
//	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&users)
//	query := "SELECT * FROM accounts WHERE email = $1"
//	row := h.DB.QueryRow(query, users.Email)
//
//	var dbUser models.Account
//	err = row.Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.SecondName, &dbUser.Email, &dbUser.Password)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			log.Println(err)
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		log.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	if users.Password != dbUser.Password {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	} else {
//		var bill = models.Bill{
//			ID:           dbUser.ID,
//			Number:       randomNumberBill(),
//			Limit:        0,
//			Cards:        nil,
//			IsBillActive: true,
//		}
//
//		query = "INSERT INTO bills (account_id, number, card, sum_limit) VALUES ($1, $2, $3, $4)"
//		_, err = h.DB.Exec(query, bill.ID, bill.Number, nil, bill.Limit)
//		if err != nil {
//			log.Println(err)
//		}
//
//		w.Header().Add("Content-Type", "application/json")
//		w.WriteHeader(http.StatusCreated)
//		json.NewEncoder(w).Encode(bill)
//
//	}
//
//}
