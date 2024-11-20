package event

import (
	"fmt"

	"github.com/syafiqparadisam/mytelebot/entity"
)

func (e *event) describing() {
	mesg := fmt.Sprintln("What are you trying to build ? \n\n Please describe your imagination ")
	e.Send(mesg)

	

	message := &entity.MessagePayload{Message: "/describing", ChatId: e.chatId}
	if err := e.repo.InsertUserCommand(message); err != nil {
		panic(err)
	}
}

