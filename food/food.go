package food

import "cookbook/model"

type Store interface {
	GetByID(int) (*model.Food, error)
	CreateFood(*model.Food) error
	UpdateFood(*model.Food) error
	DeleteFood(*model.Food) error
	List(page, limit int) ([]model.Food, int, error)
}
