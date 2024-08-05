package repositories

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
	"github.com/syafiqparadisam/mytelebot/entity"
)

type RepoInterface interface {
	CreateUser(username string) ([]Result, error)
	FindUser(username string) ([]entity.User, error)
}

type Repository struct {
	db *supabase.Client
}

func NewRepository(db *supabase.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUser(username string) ([]entity.User, error) {
	user := new([]entity.User)
	_, err := r.db.From("users").Select("username", "1", false).Eq("username", username).ExecuteTo(user)
	if err != nil {
		return nil, err
	}

	return *user, nil
}


func (r *Repository) CreateUser(username string) ([]Result, error) {
	user := struct {
		Username string `json:"username"`
	}{Username: username}

	result := []Result{}

	_, err := r.db.From("users").Insert(user, false, "", "", "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}
