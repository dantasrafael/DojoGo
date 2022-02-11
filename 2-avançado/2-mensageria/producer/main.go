package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("PRODUCER")
	topic := "arn:aws:sns:us-east-1:000000000000:GOJO_USER_CREATE"
	data := createMessage()
	sess := createSession()
	svc := sns.New(sess)
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(data),
		TopicArn: aws.String(topic),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Message ID:", *result.MessageId)
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

type Message struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Date time.Time `json:"date,omitempty"`
}

func createMessage() string {
	m := &Message{Id: uuid.New(), Date: time.Now()}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
