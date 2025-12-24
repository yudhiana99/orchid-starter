package model

type Countries struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	PhonePrefix string `json:"phone_prefix"`
	Name        string `json:"name"`
	Rank        int    `json:"rank"`
	Status      int8   `json:"status"`
	Featured    int8   `json:"featured"`
}
