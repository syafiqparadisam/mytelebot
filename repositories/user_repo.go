package repositories

import (
	"fmt"

	"github.com/syafiqparadisam/mytelebot/entity"
)

func (r *Repository) FindUser(username string) (*[]entity.User, error) {
	user := &[]entity.User{}

	_, err := r.db.From("users").Select("*", "1", false).Eq("username", username).ExecuteTo(user)
	if err != nil {
		return nil, err
	}

	fmt.Println(err)
	return user, nil
}


func (r *Repository) FindUserByChatId(chatid int64) (*[]entity.User, error) {
	user := &[]entity.User{}

	chatId := fmt.Sprintf("%d", chatid)
	_, err := r.db.From("users").Select("*", "1", false).Eq("chat_id", chatId).ExecuteTo(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (r *Repository) CreateUser(user *entity.User) error {

	_, _, err := r.db.From("users").Insert(user, false, "", "", "1").Execute()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePhone(chatid int64, phone string) error {
	phoneStruct := struct {
		PhoneNumber string `json:"phonenumber"`
	}{PhoneNumber: phone}

	// Mengonversi int64 ke string menggunakan fmt.Sprintf
	chatId := fmt.Sprintf("%d", chatid)
	_, _, err := r.db.From("users").Update(phoneStruct, "", "").Eq("chat_id", chatId).Execute()
	if err != nil {
		return err
	}
	return nil
}
