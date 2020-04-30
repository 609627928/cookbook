package store

import (
	"cookbook/model"
	"github.com/jinzhu/gorm"
)

type FoodStore struct {
	db *gorm.DB
}

func NewFoodStore(db *gorm.DB) *FoodStore {
	return &FoodStore{
		db: db,
	}
}

func (fs *FoodStore) GetByID(id int) (*model.Food, error) {
	var m model.Food
	if err := fs.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (fs *FoodStore) CreateFood(f *model.Food) (err error) {
	return fs.db.Create(f).Error
}

func (fs *FoodStore) UpdateFood(f *model.Food) error {
	return fs.db.Debug().Model(f).Update(f).Error
}

func (fs *FoodStore) DeleteFood(f *model.Food) error {
	return fs.db.Unscoped().Delete(f).Error
}

func (fs *FoodStore) List(page, limit int) ([]model.Food, int, error) {
	var (
		foods []model.Food
		count int
	)
	fs.db.Model(&foods).Count(&count)
	fs.db.Offset(page).Limit(limit).Order("-created_at").Find(&foods)
	return foods, count, nil
}
