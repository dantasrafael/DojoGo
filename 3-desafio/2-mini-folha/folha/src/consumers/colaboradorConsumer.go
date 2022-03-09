package consumers

import (
	"fmt"
	"modulo-folha/src/repositories"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const QUEUE_NAME = "COLABORADORES_FOLHA"

func IniciarColaboradorConsumer(ses *session.Session) {
	fmt.Println("Iniciando o consumidor de colaboradores...")
	canal := messaging.CreateConsumer(ses, QUEUE_NAME)

	go func() {
		for {
			select {
			case msg := <-canal:
				tratarMessage(msg)
			}
		}
	}()
}

func tratarMessage(message *messaging.ProviderMessage) {
	if message.Action == "CREATE" {
		if err := repositories.Salvar(message.Message); err != nil {
			fmt.Println(err)
		}
	} else if message.Action == "DELETE" {
		if err := repositories.Excluir(message.Message); err != nil {
			fmt.Println(err)
		}
	}
}
