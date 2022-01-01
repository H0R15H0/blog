package repository

import (
	"api/domain/model"
)

type PostRepository interface {
	FindByID(id int) (*model.Post, error)
	Create(post *model.Post) error
	Update(post *model.Post) error
	Delete(post *model.Post) error
	FindAll() ([]*model.Post, error)
}
