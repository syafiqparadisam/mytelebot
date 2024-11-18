package repositories

import "github.com/syafiqparadisam/mytelebot/entity"

func (r *Repository) GetOs() (*[]entity.Os, error) {
	oss := &[]entity.Os{}
	_, err := r.db.From("os").Select("*", "", false).ExecuteTo(oss)
	if err != nil {
		return nil, err
	}
	return oss, nil
}

func (r *Repository) GetOsByDistro(distro string) (*[]entity.Os, error) {
	os := &[]entity.Os{}
	_, err := r.db.From("os").Select("*", "", false).Eq("distro", distro).ExecuteTo(os)
	if err != nil {
		return nil, err
	}
	return os, nil
}
