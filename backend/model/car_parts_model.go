package model

type CarPart struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Part interface {
	getPartDetail() CarPart
}
