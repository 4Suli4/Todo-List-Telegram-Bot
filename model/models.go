package model

type Todo struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Done  int    `json:"done"`
}
