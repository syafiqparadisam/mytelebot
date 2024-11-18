package repositories

import (
	"github.com/supabase-community/postgrest-go"
	"github.com/syafiqparadisam/mytelebot/entity"
)

func (r *Repository) GetLastMessage(chatId int64) (*[]entity.Message, error) {
	msg := &[]entity.Message{}
	_, err := r.db.From("messages").Select("*", "1", false).Order("created_at", &postgrest.OrderOpts{Ascending: false}).ExecuteTo(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *Repository) InsertUserCommand(message *entity.MessagePayload) error {
	_, _, err := r.db.From("messages").Insert(message, false, "", "id", "1").Execute()

	if err != nil {
		return err
	}

	return nil
}
