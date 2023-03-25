package handlers

import (
	"fmt"
	"log"

	"github.com/luanpontes100/iceStorage/pkg/models"
)

func CreateTicket(t *models.Ticket) {
	fmt.Println("Digite o nome do comprador: ")
	_, err := fmt.Scanln(&t.Owner)
	if err != nil {
		log.Println("Valor passado não é válido!")
		log.Fatal(err)
	}
	fmt.Println("Digite 1 se pessoa paga meia ou 2 se a pessoa paga inteira: ")
	_, err = fmt.Scanln(&t.TicketType)
	if err != nil {
		log.Println("Valor passado não é válido!")
		log.Fatal(err)
	}
	switch t.TicketType {
	case models.Half:
		t.Price = 45.00
	case models.Full:
		t.Price = 90.00
	default:
		log.Fatal("Opção não válida")
	}

	fmt.Printf("Comprador: %s\nTipo: %s\nValor: %.2f\n", t.Owner, models.TypeTranslate[t.TicketType], t.Price)
}
