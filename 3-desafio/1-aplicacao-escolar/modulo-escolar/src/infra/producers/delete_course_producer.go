package producers

import (
	"log"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	"time"
)

const (
	ACTION_DELETE_COURSE = "DELETE_COURSE"
	TOPIC_DELETE_COURSE  = "SCHOOL_COURSE"
)

func DeleteCourseProducer(model *entities.Course) {
	message := utils.ProviderMessage{Date: time.Now(), Action: ACTION_DELETE_COURSE, Message: *model}

	log.Printf("Enviando msg para o tópico '%s': %v\n", TOPIC_DELETE_COURSE, message.GetJsonMessage())
}
