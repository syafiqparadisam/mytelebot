package repositories

import "github.com/syafiqparadisam/mytelebot/entity"

func (r *Repository) InsertOrder(entity *entity.Order) error {
	_, _, err := r.db.From("order").Insert(entity, false, "", "", "").Execute()
	if err != nil {
		return err
	}
	return nil
}
