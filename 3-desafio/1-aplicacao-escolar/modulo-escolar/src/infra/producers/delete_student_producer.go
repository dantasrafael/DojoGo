package producers

import (
	"encoding/json"
	"log"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_DELETE_STUDENT = "DELETE_STUDENT"
	TOPIC_DELETE_STUDENT  = "SCHOOL_STUDENT"
)

func DeleteStudentProducer(model *entities.Student) {
	modelJson, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	message := messaging.ProviderMessage{Action: ACTION_DELETE_STUDENT, Message: string(modelJson)}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	if err = messaging.PublishMessage(svc, TOPIC_DELETE_STUDENT, message.GetJson()); err != nil {
		log.Printf("could not send message topic %s: %v", TOPIC_CREATE_ENROLLMENT, err)
	}
}
