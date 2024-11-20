package repositories

import (
	"fmt"

	"github.com/syafiqparadisam/mytelebot/entity"
)

func (r *Repository) InsertApp(entity *entity.AppPayload) error {
	_, _, err := r.db.From("application").Insert(entity, false, "", "", "").Execute()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTechUsed(entity *entity.UpdateTech) error {
	tech := struct {
		Tech string `json:"tech"`
	}{Tech: entity.Tech}

	_, _, err := r.db.From("application").Update(tech, "", "").Eq("id", fmt.Sprint(entity.Id)).Execute()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetApp(id int) (*[]entity.App, error) {
	entity := &[]entity.App{}

	_, err := r.db.From("application").Select("*", "", false).Eq("id", fmt.Sprint("%d", id)).ExecuteTo(entity)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
