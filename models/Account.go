package models

type Account struct {
	ID         int    `json:"ID"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bill       []Bill `json:"bills"`
}
