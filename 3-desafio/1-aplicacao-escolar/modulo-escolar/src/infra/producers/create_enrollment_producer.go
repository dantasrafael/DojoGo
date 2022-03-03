package producers

import (
	"encoding/json"
	"log"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_CREATE_ENROLLMENT = "CREATE_ENROLLMENT"
	TOPIC_CREATE_ENROLLMENT  = "SCHOOL_ENROLLMENT"
)

func CreateEnrollmentProducer(model *entities.Enrollment) {
	modelJson, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	message := messaging.ProviderMessage{Action: ACTION_CREATE_ENROLLMENT, Message: string(modelJson)}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	if err = messaging.PublishMessage(svc, TOPIC_CREATE_ENROLLMENT, message.GetJson()); err != nil {
		log.Printf("could not send message topic %s: %v", TOPIC_CREATE_ENROLLMENT, err)
	}
}
