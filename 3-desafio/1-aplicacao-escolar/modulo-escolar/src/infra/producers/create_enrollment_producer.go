package producers

import (
	"log"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	"time"
)

const (
	ACTION_CREATE_ENROLLMENT = "CREATE_ENROLLMENT"
	TOPIC_CREATE_ENROLLMENT  = "SCHOOL_ENROLLMENT"
)

func CreateEnrollmentProducer(model *entities.Enrollment) {
	message := utils.ProviderMessage{Date: time.Now(), Action: ACTION_CREATE_ENROLLMENT, Message: *model}

	log.Printf("Enviando msg para o tópico '%s': %v\n", TOPIC_CREATE_ENROLLMENT, message.GetJsonMessage())
}
