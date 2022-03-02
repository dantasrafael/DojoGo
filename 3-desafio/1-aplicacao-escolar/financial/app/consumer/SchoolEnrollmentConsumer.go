package consumer

import (
	"encoding/json"
	"financial/app/consumer/model"
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

const schoolEnrollmentQueueName = "SCHOOL_ENROLLMENT_FINANCIAL"

func StartSchoolEnrollmentConsumer(sess *session.Session) {
	ch := createConsumer(sess)

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}

func createConsumer(sess *session.Session) chan *model.Enrollment {
	ch := make(chan *model.Enrollment, 1)

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
				var n model.Enrollment
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
