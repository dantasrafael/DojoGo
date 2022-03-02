package producers

import (
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_DELETE_ENROLLMENT = "DELETE_ENROLLMENT"
	TOPIC_DELETE_ENROLLMENT  = "SCHOOL_ENROLLMENT"
)

func DeleteEnrollmentProducer(model *entities.Enrollment) {
	message := utils.ProviderMessage{Action: ACTION_DELETE_ENROLLMENT, Message: *model}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	messaging.PublishMessage(svc, TOPIC_DELETE_ENROLLMENT, message.GetJsonMessage())
}
