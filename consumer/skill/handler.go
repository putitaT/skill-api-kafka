package skill

import (
	"fmt"

	"github.com/IBM/sarama"
)

func HandleMessage(msg *sarama.ConsumerMessage, session sarama.ConsumerGroupSession) {
	var err error
	switch string(msg.Key) {
	case "create":
		err = CreateSkill(msg.Value)
	case "update":
		err = updateSkillById(msg.Value)
	case "updateName":
		err = updateNameById(msg.Value)
	case "updateDesc":
		err = updateDescriptionById(msg.Value)
	case "updateLogo":
		err = updateLogoById(msg.Value)
	case "updateTags":
		err = updateTagById(msg.Value)
	case "delete":
		err = deleteSkillById(msg.Value)
	}

	if err == nil {
		session.MarkMessage(msg, "")
	} else {
		fmt.Println(err)
	}
}
