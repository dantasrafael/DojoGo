package producers

import (
	"encoding/json"
	"log"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_DELETE_COURSE = "DELETE_COURSE"
	TOPIC_DELETE_COURSE  = "SCHOOL_COURSE"
)

func DeleteCourseProducer(model *entities.Course) {
	modelJson, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	message := messaging.ProviderMessage{Action: ACTION_DELETE_COURSE, Message: string(modelJson)}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	if err = messaging.PublishMessage(svc, TOPIC_DELETE_COURSE, message.GetJson()); err != nil {
		log.Printf("could not send message topic %s: %v", TOPIC_CREATE_ENROLLMENT, err)
	}
}
