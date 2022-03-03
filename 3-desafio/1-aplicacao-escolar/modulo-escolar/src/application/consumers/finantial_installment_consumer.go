package consumers

import (
	"log"
	usecases "modulo-escolar/src/domain/usecases/enrollment"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const queueName = "FINANCIAL_INSTALLMENT_SCHOOL"

func StartFinantialInstallmentConsumer(sess *session.Session) {
	ch := messaging.CreateConsumer(sess, queueName)

	go func() {
		for {
			select {
			case msg := <-ch:
				if err := usecases.UpdateEnrollmentStatusUsecase(msg.Message); err != nil {
					log.Printf("ERROR: could not consume message: \n%v: \n%v", msg, err)
				}
			}
		}
	}()
}
