package producers

import (
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_DELETE_STUDENT = "DELETE_STUDENT"
	TOPIC_DELETE_STUDENT  = "SCHOOL_STUDENT"
)

func DeleteStudentProducer(model *entities.Student) {
	message := utils.ProviderMessage{Action: ACTION_DELETE_STUDENT, Message: *model}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	messaging.PublishMessage(svc, TOPIC_DELETE_STUDENT, message.GetJsonMessage())
}
