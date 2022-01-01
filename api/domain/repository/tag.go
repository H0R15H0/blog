package repository

import (
	"api/domain/model"
)

type TagRepository interface {
	FindByID(id int) (*model.Tag, error)
	Create(tag *model.Tag) error
	Update(tag *model.Tag) error
	Delete(tag *model.Tag) error
	FindAll() ([]*model.Tag, error)
}
