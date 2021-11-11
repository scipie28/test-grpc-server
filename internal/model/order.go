package model

type Order struct {
	Id          string
	Items       []string
	Description string
	Price       float32
	Destination string
}
