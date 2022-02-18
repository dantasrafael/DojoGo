package producers

import (
	"log"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	"time"
)

const (
	ACTION_DELETE_ENROLLMENT = "DELETE_ENROLLMENT"
	TOPIC_DELETE_ENROLLMENT  = "SCHOOL_ENROLLMENT"
)

func DeleteEnrollmentProducer(model *entities.Enrollment) {
	message := utils.ProviderMessage{Date: time.Now(), Action: ACTION_DELETE_ENROLLMENT, Message: *model}

	log.Printf("Enviando msg para o tópico '%s': %v\n", TOPIC_DELETE_ENROLLMENT, message.GetJsonMessage())
}
