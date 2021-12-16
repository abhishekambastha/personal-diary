package domain

type Entry struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Text    string `json:"text"`
	Picture string `json:"picture"`
	Tid     string `json:"tid"`
}
