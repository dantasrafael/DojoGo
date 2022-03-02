package producers

import (
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
)

const (
	ACTION_DELETE_COURSE = "DELETE_COURSE"
	TOPIC_DELETE_COURSE  = "SCHOOL_COURSE"
)

func DeleteCourseProducer(model *entities.Course) {
	message := utils.ProviderMessage{Action: ACTION_DELETE_COURSE, Message: *model}
	sess := messaging.CreateLocalstackSession()
	svc := sns.New(sess)

	messaging.PublishMessage(svc, TOPIC_DELETE_COURSE, message.GetJsonMessage())
}
