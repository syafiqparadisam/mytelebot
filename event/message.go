package event

import (
	"fmt"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/syafiqparadisam/mytelebot/entity"
	"github.com/syafiqparadisam/mytelebot/utils"
)

func (e *event) handleMessage(update tgbotapi.Update, user string) {
	chat := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	message := &entity.MessagePayload{Message: chat.Text, ChatId: chat.ChatID}
	if err := e.repo.InsertUserCommand(message); err != nil {
		panic(err)
	}

	var msg string
	switch chat.Text {
	case "/start":
		e.sendWelcomeMessage(update, user)
		content := utils.ReadFile("services.txt")
		e.Send(chat.ChatID, content)
	case "/1":
		msg = "Before that, please describe your website"
		e.Send(chat.ChatID, msg)
	case "/2":
		e.GetOs(chat)
	default:
		lastMsgs, err := e.repo.GetLastMessage(chat.ChatID)
		if err != nil {
			panic(err)
		}

		mesg := *lastMsgs

		if mesg[1].Message == "/2" {
			e.handleChooseDistro(chat)
		} else if mesg[1].Message == "/phonenumber" {
			isValid := e.validatePhoneNumber(chat.Text)
			if !isValid {
				msg := fmt.Sprintln("Please enter the right phone number")
				e.Send(chat.ChatID, msg)
				message := &entity.MessagePayload{Message: "/phonenumber", ChatId: chat.ChatID}
				if err := e.repo.InsertUserCommand(message); err != nil {
					panic(err)
				}
				return
			}

			if err := e.repo.UpdatePhone(chat.ChatID, chat.Text); err != nil {
				panic(err)
			}

			msg = fmt.Sprintln("Alright, my boss will call you later \n\n Thanks for buy my services, we will hope you enjoy our services")
			e.Send(chat.ChatID, msg)

			e.sendWelcomeMessage(update, user)
			content := utils.ReadFile("services.txt")
			e.Send(chat.ChatID, content)
			return
		} else {

			content := utils.ReadFile("default.txt")
			e.Send(chat.ChatID, content)
		}
	}

}

func (e *event) GetOs(chat tgbotapi.MessageConfig) {
	result, err := e.repo.GetOs()
	if err != nil {
		panic(err)
	}

	var osFormatted string
	for i, os := range *result {
		osFormatted += fmt.Sprintf("%d. %s %s \n", i+1, os.Distro, os.Level)
	}

	msg := fmt.Sprintf("What types of linux do you want ? \n%v\n\n For example (arch, debian)", osFormatted)
	e.Send(chat.ChatID, msg)
}

func (e *event) validatePhoneNumber(phone string) bool {
	// Fungsi untuk validasi nomor telepon
	// Membuat regex yang memeriksa kondisi
	// Dimulai dengan 0, hanya angka, dan panjang antara 11 hingga 13 digit
	re := regexp.MustCompile(`^0\d{10,12}$`)

	// Mengecek apakah nomor telepon sesuai dengan pola regex
	return re.MatchString(phone)
}

func (e *event) sendWelcomeMessage(update tgbotapi.Update, user string) {
	content := utils.ReadFile("welcome.txt")
	message := fmt.Sprintf("Welcome back %s \n %s", user, content)
	e.Send(update.Message.Chat.ID, message)
}

func (e *event) handleChooseDistro(chat tgbotapi.MessageConfig) {
	result, err := e.repo.GetOs()
	if err != nil {
		panic(err)
	}
	var osFormatted string
	for _, os := range *result {
		osFormatted += os.Distro
	}

	if !strings.Contains(osFormatted, chat.Text) {
		msg := fmt.Sprintf("Operating system %s not found, please type the right operating system in the lists ", chat.Text)
		e.Send(chat.ChatID, msg)
		e.GetOs(chat)
		message := &entity.MessagePayload{Message: "/2", ChatId: chat.ChatID}
		if err := e.repo.InsertUserCommand(message); err != nil {
			panic(err)
		}
	} else {
		msg := fmt.Sprintf("Alright my boss will install you an %s operating system", chat.Text)
		e.Send(chat.ChatID, msg)

		result, err := e.repo.GetOsByDistro(chat.Text)

		if err != nil {
			panic(err)
		}
		os := *result

		price := e.CalculatePriceLevelDistro(os[0])
		msg = fmt.Sprintf("Price of %s is Rp. %d \n", os[0].Distro, price)
		e.Send(chat.ChatID, msg)

		order := &entity.Order{ServicesType: "os", ServicesId: os[0].Id, ChatId: chat.ChatID, Price: price}
		err = e.repo.InsertOrder(order)
		if err != nil {
			panic(err)
		}

		askPhone := fmt.Sprintln("Please type your phone number, so myboss can chat you later")
		e.Send(chat.ChatID, askPhone)

		message := &entity.MessagePayload{Message: "/phonenumber", ChatId: chat.ChatID}
		if err := e.repo.InsertUserCommand(message); err != nil {
			panic(err)
		}
	}
}

func (e *event) WantOrder(chat tgbotapi.MessageConfig, order *entity.Order, price int64) {
	err := e.repo.InsertOrder(order)
	if err != nil {
		panic(err)
	}

	users, err := e.repo.FindUserByChatId(chat.ChatID)
	if err != nil {
		panic(err)
	}
	askPhone := fmt.Sprintln("Please type your phone number, so myboss can chat you later")
	e.Send(chat.ChatID, askPhone)

	message := &entity.MessagePayload{Message: "/phonenumber", ChatId: chat.ChatID}
	if err := e.repo.InsertUserCommand(message); err != nil {
		panic(err)
	}
}
