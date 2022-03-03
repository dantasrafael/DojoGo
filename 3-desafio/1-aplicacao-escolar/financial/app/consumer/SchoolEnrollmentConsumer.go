package consumer

import (
	"context"
	"financial/app/consumer/model"
	"financial/domain/entity"
	"financial/domain/usecase"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const schoolEnrollmentQueueName = "SCHOOL_ENROLLMENT_FINANCIAL"

func StartSchoolEnrollmentConsumer(sess *session.Session) {
	ch := messaging.CreateConsumer(sess, schoolEnrollmentQueueName)
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
