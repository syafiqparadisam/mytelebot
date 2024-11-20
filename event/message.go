package event

import (
	"fmt"

	"github.com/syafiqparadisam/mytelebot/entity"
	"github.com/syafiqparadisam/mytelebot/utils"
)

func (e *event) handleMessage() {

	message := &entity.MessagePayload{Message: e.chat, ChatId: e.chatId}
	if err := e.repo.InsertUserCommand(message); err != nil {
		panic(err)
	}

	var msg string
	switch e.chat {
	case "/start":
		e.start()
	case "/1":
		e.describing()
	case "/2":
		e.GetOs()
	case "/3":
		e.Send("sek belum")
	case "/4":
		e.describing()
	default:
		lastMsgs, err := e.repo.GetLastMessage(e.chatId)
		if err != nil {
			panic(err)
		}

		mesg := *lastMsgs

		if mesg[1].Message == "/2" {
			e.handleChooseDistro()
		} else if mesg[1].Message == "/phonenumber" {
			isValid := e.validatePhoneNumber()
			if !isValid {
				msg := fmt.Sprintln("Please enter the right phone number")
				e.Send(msg)

				e.insertCommand("/phonenumber")
				return
			}

			if err := e.repo.UpdatePhone(e.chatId, e.chat); err != nil {
				panic(err)
			}

			e.Done()
			return
		} else if mesg[1].Message == "/confirmos" {
			if e.chat == "/yes" {

				osChoose := mesg[2].Message
				result, err := e.repo.GetOsByDistro(osChoose)

				if err != nil {
					panic(err)
				}
				os := *result
				price := os[0].Price
				servicesId := os[0].Id
				order := &entity.Order{ChatId: e.chatId, Price: &price, ServicesType: "os", ServicesId: &servicesId}
				e.WantOrder(order)
			} else {
				msg := fmt.Sprintln("Canceling order")
				e.Send(msg)
				e.start()
			}
		} else if mesg[1].Message == "/describing" {
			isValid := e.ensureIsNotCommandAndFullNumber()

			if isValid {
				app := &entity.AppPayload{Description: e.chat}
				if err := e.repo.InsertApp(app); err != nil {
					panic(err)
				}

				msg = fmt.Sprintln("Wow amazing project, i think its awesome \n What tech did you want to use in this project ?")
				e.Send(msg)

				e.insertCommand("/techuse")
				return
			}

			msg = fmt.Sprintln("Please desribe your website properly")
			e.Send(msg)

			e.insertCommand("/describing")
		} else if mesg[1].Message == "/techuse" {

			isValid := e.ensureIsNotCommandAndFullNumber()

			if isValid {
				update := &entity.UpdateTech{Tech: e.chat, Description: mesg[2].Message}
				if err := e.repo.UpdateTechUsed(update); err != nil {
					panic(err)
				}

				apps, err := e.repo.GetApp(mesg[2].Message, e.chat)
				if err != nil {
					panic(err)
				}
				app := *apps
				servicesId := app[0].Id

				order := &entity.Order{ChatId: e.chatId, Price: nil, ServicesType: "app", ServicesId: &servicesId}
				e.WantOrder(order)
				return
			}

			msg = fmt.Sprintln("Please tell the right tech")
			e.Send(msg)

			e.insertCommand("/techuse")
		} else {

			content := utils.ReadFile("default.txt")
			e.Send(content)
		}
	}

}
