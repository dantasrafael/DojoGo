package messaging

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsNotification struct {
	MessageId string `json:"MessageId"`
	Message   string `json:"Message"`
}

func CreateConsumer(sess *session.Session, queueName string) chan *ProviderMessage {
	ch := make(chan *ProviderMessage, 1)

	svc := sqs.New(sess)
	queueResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if err != nil {
		log.Printf("could not connect to queue %s: %v\n", queueName, err)
		os.Exit(1)
	}

	go func() {
		for {
			msgs, err := readMessages(svc, queueResult)
			if err != nil {
				log.Printf("could not receive message: %v\n", err)
				continue
			}
			if len(msgs.Messages) > 0 {
				msg := msgs.Messages[0]
				var n sqsNotification
				err = json.Unmarshal([]byte(*msg.Body), &n)
				if err != nil {
					log.Printf("Could not unmarshal message body [%s]: %v\n", *msg.Body, err)
				}
				var pm ProviderMessage
				err = json.Unmarshal(([]byte(n.Message)), &pm)
				if err != nil {
					log.Printf("Could not unmarshal message body [%s]: %v\n", *msg.Body, err)
				} else {
					ch <- &pm
					removeMessageFromQueue(svc, queueResult, msg)
				}
			} else {
				log.Println("No messages in queue..")
			}
		}
	}()

	return ch
}

func readMessages(svc *sqs.SQS, queueResult *sqs.GetQueueUrlOutput) (*sqs.ReceiveMessageOutput, error) {
	var msgs, err = svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:              queueResult.QueueUrl,
		MaxNumberOfMessages:   aws.Int64(1),
		WaitTimeSeconds:       aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
	})
	return msgs, err
}

func removeMessageFromQueue(svc *sqs.SQS, queueResult *sqs.GetQueueUrlOutput, msg *sqs.Message) {
	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueResult.QueueUrl,
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		log.Printf("Could not delete message %s: %v", msg, err)
	}
}
