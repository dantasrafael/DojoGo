package producers

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"log"
	"modulo-colaboradores/model"
)

const (
	NOME_TOPICO  = "COLABORADORES"
	ACAO_CRIACAO = "CREATE"
	ACAO_DELECAO = "DELETE"
)

func CriacaoColaborador(colaborador *model.Colaborador) {
	modelJson, err := json.Marshal(colaborador)
	if err != nil {
		log.Fatal(err)
	}
	message := messaging.ProviderMessage{
		Action:  ACAO_CRIACAO,
		Message: string(modelJson),
	}
	sessao := messaging.CreateLocalstackSession()
	service := sns.New(sessao)
	err = messaging.PublishMessage(service, NOME_TOPICO, message.GetJson())
	if err != nil {
		log.Fatal(err)
	}
}

func DelecaoColaborador(colaborador *model.Colaborador) {
	modelJson, err := json.Marshal(colaborador)
	if err != nil {
		log.Fatal(err)
	}
	message := messaging.ProviderMessage{
		Action:  ACAO_DELECAO,
		Message: string(modelJson),
	}
	sessao := messaging.CreateLocalstackSession()
	service := sns.New(sessao)
	err = messaging.PublishMessage(service, NOME_TOPICO, message.GetJson())
	if err != nil {
		log.Fatal(err)
	}
}
