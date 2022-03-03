package consumer

import (
	"context"
	"database/sql"
	"encoding/json"
	"financial/app/consumer/model"
	"financial/domain/entity"
	"financial/domain/usecase"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const schoolEnrollmentQueueName = "SCHOOL_ENROLLMENT_FINANCIAL"

func StartSchoolEnrollmentConsumer(sess *session.Session, db *sql.DB) {
	ch := messaging.CreateConsumer(sess, schoolEnrollmentQueueName)
	uc := usecase.NewRegisterAccount(db)

	go func() {
		for {
			select {
			case msg := <-ch:
				var e model.Enrollment
				err := json.Unmarshal([]byte(msg.Message), &e)
				if err != nil {
					log.Printf("ERROR could not parse message:\n%v\n%v", msg, err)
					continue
				}
				account := &entity.Account{
					ClientID:     e.StudentID,
					CourseID:     e.CourseID,
					Installments: e.Installments,
					Total:        e.Total,
				}
				ctx := context.WithValue(context.Background(), "origin", "school-enrollment-consumer")
				log.Printf("starting message processing...\n%v\n", account)
				err = uc.Execute(ctx, account)
				if err != nil {
					log.Printf("ERROR: could not consume message: \n%v: \n%v", msg, err)
				}
			}
		}
	}()
}
