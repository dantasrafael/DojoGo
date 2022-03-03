package producers

import (
	"context"
	"encoding/json"
	"financial/domain/entity"
	"financial/infra/producers/model"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"log"
	"strconv"
)

const (
	actionRegisterAccount = "REGISTER_ACCOUNT"
	topicRegisterAccount  = "FINANCIAL_INSTALLMENT"
)

type RegisterAccountProducer struct {
}

func NewRegisterAccountProducer() RegisterAccountProducer {
	return RegisterAccountProducer{}
}

func (p RegisterAccountProducer) Send(_ context.Context, account *entity.Account) {
	externalID, _ := strconv.ParseUint(account.ExternalID, 10, 64)
	m, err := json.Marshal(model.RegisterAccountNotification{ID: externalID, Status: account.Status})
	if err != nil {
		log.Fatal(err)
	}

	message := messaging.ProviderMessage{Action: actionRegisterAccount, Message: string(m)}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	err = messaging.PublishMessage(svc, topicRegisterAccount, message.GetJson())
	if err != nil {
		log.Printf("ERROR: could not send message to topic %s: %v\n", topicRegisterAccount, err)
	}
}
