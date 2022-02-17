package messaging

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
)

func CreateLocalstackSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:4566"),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(uuid.NewString(), "no_secret", "no_token"),
	}))
}
