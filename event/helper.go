package event

import (
	"fmt"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/syafiqparadisam/mytelebot/entity"
	"github.com/syafiqparadisam/mytelebot/utils"
)

func (e *event) start() {

	welcome := utils.ReadFile("welcome.txt")
	message := fmt.Sprintf("Welcome back %s \n %s", e.user, welcome)

	e.Send(message)
	content := utils.ReadFile("services.txt")
	e.Send(content)
}

// func (e *event) CalculatePriceLevelDistro(os entity.Os) int64 {
// 	level := os.Level
// 	switch level {
// 	case "veryeasy":
// 		return 20000
// 	case "easy":
// 		return 30000
// 	case "medium":
// 		return 40000
// 	case "hard":
// 		return 60000
// 	case "veryhard":
// 		return 100000
// 	}
// 	return 0
// }

func (e *event) validatePhoneNumber() bool {
	// Fungsi untuk validasi nomor telepon
	// Membuat regex yang memeriksa kondisi
	// Dimulai dengan 0, hanya angka, dan panjang antara 11 hingga 13 digit
	re := regexp.MustCompile(`^0\d{10,12}$`)

	// Mengecek apakah nomor telepon sesuai dengan pola regex
	return re.MatchString(e.chat)
}

func (e *event) WantOrder(order *entity.Order) {
	err := e.repo.InsertOrder(order)
	if err != nil {
		panic(err)
	}

	users, err := e.repo.FindUserByChatId(e.chatId)
	if err != nil {
		panic(err)
	}

	user := *users

	if user[0].PhoneNumber == nil {

		askPhone := fmt.Sprintln("Please type your phone number, so myboss can chat you later")
		e.Send(askPhone)

		message := &entity.MessagePayload{Message: "/phonenumber", ChatId: e.chatId}
		if err := e.repo.InsertUserCommand(message); err != nil {
			panic(err)
		}
		return
	}

	e.Done()
}

func (e *event) ensureIsNotCommandAndFullNumber() bool {
	// Trim input to remove unnecessary spaces
	trimmedInput := strings.TrimSpace(e.chat)

	// Check if input starts with '/'
	if strings.HasPrefix(trimmedInput, "/") {
		return false
	}

	// Ensure the input is NOT fully numeric
	fullNumRgx := regexp.MustCompile(`^[0-9]+$`)
	return !fullNumRgx.MatchString(trimmedInput) 
}

func (e *event) Send(message string) {
	msg := tgbotapi.NewMessage(e.chatId, message)
	if _, err := e.bot.Send(msg); err != nil {
		panic(err)
	}
}

func (e *event) Done() {
	msg := fmt.Sprintln("Alright, my boss will call you later \n\n Thanks for ordering my services, we will hope you enjoy our services")
	e.Send(msg)
	e.start()
}

func (e *event) insertCommand(command string) {
	message := &entity.MessagePayload{Message: command, ChatId: e.chatId}
	if err := e.repo.InsertUserCommand(message); err != nil {
		panic(err)
	}
}
