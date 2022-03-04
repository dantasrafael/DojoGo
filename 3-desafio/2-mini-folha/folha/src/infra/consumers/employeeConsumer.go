package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type Notification struct {
	Type      string    `json:"Type,omitempty"`
	MessageId string    `json:"MessageId,omitempty"`
	Message   string    `json:"Message,omitempty"`
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

var (
	poolWorkers = 1
)

func main() {
	queueName := "EMPLOYEE_CREATE_CONSUMER"

	ch := createConsumer(queueName)

	for {
		select {
		case msg := <-ch:
			processMessage(msg)
		}
	}
}

func createConsumer(queueName string) chan *Notification {
	ch := make(chan *Notification, 1)

	sess := createSession()
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
				var n Notification
				log.Println(*msg.Body)
				err = json.Unmarshal([]byte(*msg.Body), &n)
				if err != nil {
					log.Printf("Could not unmarshal message body [%s]: %v\n", *msg.Body, err)
				} else {
					ch <- &n
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
	msgs, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:              queueResult.QueueUrl,
		MaxNumberOfMessages:   aws.Int64(1),
		WaitTimeSeconds:       aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
	})
	return msgs, err
}

func processMessage(msg *Notification) {
	log.Printf("Processing message: %v\n", msg)
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

func createSession() *session.Session {
	var sess *session.Session
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		// Initialize a session that the SDK will use to load
		// credentials from the shared credentials file. (~/.aws/credentials).
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	} else {
		sess = session.Must(session.NewSession(&aws.Config{
			Region:      aws.String("us-east-1"),
			Endpoint:    aws.String("http://localhost:4566"),
			DisableSSL:  aws.Bool(true),
			Credentials: credentials.NewStaticCredentials(uuid.NewString(), "no_secret", "no_token"),
		}))
	}
	return sess
}
