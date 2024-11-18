package repositories

import (
	"github.com/supabase-community/supabase-go"
	"github.com/syafiqparadisam/mytelebot/entity"
)

type RepoInterface interface {
	CreateUser(user *entity.User) error
	FindUser(username string) (*[]entity.User, error)
	InsertUserCommand(message *entity.MessagePayload) error
	GetOs() (*[]entity.Os, error)
	GetLastMessage(chatId int64) (*[]entity.Message, error)
	GetOsByDistro(distro string) (*[]entity.Os, error)
	InsertOrder(entity *entity.Order) error
	UpdatePhone(chatid int64, phone string) error
	FindUserByChatId(chatid int64) (*[]entity.User, error)
}

type Repository struct {
	db *supabase.Client
}

func NewRepository(db *supabase.Client) *Repository {
	return &Repository{db: db}
}
