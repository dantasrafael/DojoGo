package messaging

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

func PublishMessage(svc *sns.SNS, topic string, data string) error {
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(data),
		TopicArn: aws.String(fmt.Sprintf("arn:aws:sns:us-east-1:000000000000:%s", topic)),
	})
	if err != nil {
		return err
	}
	fmt.Println("Send message with id:", *result.MessageId)
	return nil
}
