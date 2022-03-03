package consumer

import (
	"context"
	"encoding/json"
	"financial/app/consumer/model"
	"financial/domain/entity"
	"financial/domain/usecase"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"log"
	"os"
)

const schoolEnrollmentQueueName = "SCHOOL_ENROLLMENT_FINANCIAL"

func StartSchoolEnrollmentConsumer(sess *session.Session) {
	ch := createConsumer(sess)
	uc := usecase.NewRegisterAccount()

	for {
		select {
		case msg := <-ch:
			enrollment, ok := msg.Message.(*model.Enrollment)
			if !ok {
				log.Printf("ERROR: could not parse message: \n%v", msg)
				continue
			}
			account := &entity.Account{
				ClientID:     enrollment.StudentID,
				CourseID:     enrollment.CourseID,
				Installments: enrollment.Installments,
				Total:        enrollment.Total,
			}
			ctx := context.WithValue(context.Background(), "origin", "school-enrollment-consumer")
			err := uc.Execute(ctx, account)
			if err != nil {
				log.Printf("ERROR: could not consume message: \n%v: \n%v", msg, err)
			}
		}
	}
}

func createConsumer(sess *session.Session) chan *messaging.ProviderMessage {
	ch := make(chan *messaging.ProviderMessage, 1)

	svc := sqs.New(sess)
	queueResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(schoolEnrollmentQueueName)})
	if err != nil {
		log.Printf("could not connect to queue %s: %v\n", schoolEnrollmentQueueName, err)
		os.Exit(1)
	}

	go func() {
		for {
			msgs, err := messaging.ReadMessages(svc, queueResult)
			if err != nil {
				log.Printf("could not receive message: %v\n", err)
				continue
			}
			if len(msgs.Messages) > 0 {
				msg := msgs.Messages[0]
				var n messaging.ProviderMessage
				log.Println(*msg.Body)
				err = json.Unmarshal([]byte(*msg.Body), &n)
				if err != nil {
					log.Printf("Could not unmarshal message body [%s]: %v\n", *msg.Body, err)
				} else {
					ch <- &n
					messaging.RemoveMessageFromQueue(svc, queueResult, msg)
				}
			} else {
				log.Println("No messages in queue..")
			}
		}
	}()

	return ch
}
