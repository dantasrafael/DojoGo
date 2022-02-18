package producers

import (
	"log"
	"modulo-escolar/src/core/utils"
	"modulo-escolar/src/domain/entities"
	"time"
)

const (
	ACTION_DELETE_STUDENT = "DELETE_STUDENT"
	TOPIC_DELETE_STUDENT  = "SCHOOL_STUDENT"
)

func DeleteStudentProducer(model *entities.Student) {
	message := utils.ProviderMessage{Date: time.Now(), Action: ACTION_DELETE_STUDENT, Message: *model}

	log.Printf("Enviando msg para o tópico '%s': %v\n", TOPIC_DELETE_STUDENT, message.GetJsonMessage())
}
