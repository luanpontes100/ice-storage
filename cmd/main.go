package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/luanpontes100/iceStorage/pkg/handlers"
	"github.com/luanpontes100/iceStorage/pkg/models"
	"github.com/luanpontes100/iceStorage/pkg/utils"
)

func main() {
	var tickets []models.Ticket
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		// Waits for SIGINT or SIGTERM
		<-sigChan
		utils.SaveTickets(tickets)
		os.Exit(0)
	}()

	var choice int
	utils.LoadTickets(&tickets)
	fmt.Println("Bem vindo ao portal de gerenciamento da Iceland")
	for {
		fmt.Println("Tickets atuais: ", tickets)
		fmt.Println("Selecione uma opção:")
		fmt.Printf("1) Novo Ingresso\n2) Atualizar Ingresso\n3) Pesquisar Ingresso\n4) Deletar Ingresso\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			log.Println("Valor passado não é válido!")
			log.Fatal(err)
		}
		switch choice {
		case 1:
			var ticket models.Ticket
			handlers.CreateTicket(&ticket)
			tickets = append(tickets, ticket)
		case 2:

		default:
			fmt.Println("Programa encerrado")
		}
	}
}
