package producers

import (
	"context"
	"financial/domain/entity"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"log"
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

func (p RegisterAccountProducer) Send(ctx context.Context, account *entity.Account) {

	message := messaging.ProviderMessage{Action: actionRegisterAccount, Message: *account}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	err := messaging.PublishMessage(svc, topicRegisterAccount, message.GetJson())
	if err != nil {
		log.Printf("ERROR: could not send message to topic %s: %v\n", topicRegisterAccount, err)
	}
}