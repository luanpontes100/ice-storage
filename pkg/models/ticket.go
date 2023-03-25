package models

const Half, Full int = 1, 2

var TypeTranslate = map[int]string{
	1: "Meia",
	2: "Inteira",
}

type Ticket struct {
	Owner      string
	TicketType int
	Price      float32
}
