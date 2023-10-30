package handlers

//
//import (
//	"bank-di/models"
//	"database/sql"
//	"encoding/json"
//	"io"
//	"log"
//	"net/http"
//)
//
//func (h Handler) CloseBillAndCards(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	var data struct {
//		AccountID int    `json:"account_id"`
//		Password  string `json:"password"`
//		BillID    int    `json:"bill_id"`
//	}
//	err = json.Unmarshal(body, &data)
//	if err != nil {
//		log.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	// Проверка существования аккаунта и счета в базе данных
//	query := "SELECT * FROM accounts WHERE id = $1"
//	row := h.DB.QueryRow(query, data.AccountID)
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
//
//	// Проверка соответствия пароля
//	if data.Password != dbUser.Password {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//
//	// Закрытие счета
//	query = "UPDATE bills SET is_bill_active = false WHERE id = $1 AND account_id = $2"
//	_, err = h.DB.Exec(query, data.BillID, data.AccountID)
//	if err != nil {
//		log.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	// Закрытие всех привязанных карт
//	query = "UPDATE cards SET is_card_active = false WHERE bill_id = $1"
//	_, err = h.DB.Exec(query, data.BillID)
//	if err != nil {
//		log.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
