package event

import (
	"fmt"
	"strings"

	"github.com/syafiqparadisam/mytelebot/entity"
)

func (e *event) GetOs() {
	result, err := e.repo.GetOs()
	if err != nil {
		panic(err)
	}

	var osFormatted string
	for i, os := range *result {
		osFormatted += fmt.Sprintf("%d. %s Rp. %d \n", i+1, os.Distro, os.Price)
	}

	msg := fmt.Sprintf("What types of linux do you want ? \n%v\n\n For example (arch, debian)", osFormatted)
	e.Send(msg)
}

func (e *event) handleChooseDistro() {
	result, err := e.repo.GetOs()
	if err != nil {
		panic(err)
	}
	var osFormatted string
	for _, os := range *result {
		osFormatted += os.Distro
	}

	if !strings.Contains(osFormatted, e.chat) {
		msg := fmt.Sprintf("Operating system %s not found, please type the right operating system in the lists ", e.chat)
		e.Send(msg)
		e.GetOs()

		message := &entity.MessagePayload{Message: "/2", ChatId: e.chatId}
		if err := e.repo.InsertUserCommand(message); err != nil {
			panic(err)
		}
	} else {
		msg := fmt.Sprintf("Alright my boss will install you an %s operating system", e.chat)
		e.Send(msg)

		msg = fmt.Sprintf("Are you sure you choose this %s operating system ? \n\n type /yes for confirm and /no for canceling", e.chat)
		e.Send(msg)

		message := &entity.MessagePayload{Message: "/confirmos", ChatId: e.chatId}
		if err := e.repo.InsertUserCommand(message); err != nil {
			panic(err)
		}

	}
}
