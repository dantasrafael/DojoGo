package messaging

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/google/uuid"
)

func CreateLocalstackSession() *session.Session {
	host := config.LocalStackHost
	return session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String(fmt.Sprintf("http://%s:4566", host)),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(uuid.NewString(), "no_secret", "no_token"),
	}))
}
