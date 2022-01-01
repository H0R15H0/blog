package usecase

import (
	"api/domain/model"
	"api/domain/repository"
)

type PostParamsObject model.Post

// PostUsecase post usecaseのinterface
type PostUsecase interface {
	Create(post PostParamsObject) (*model.Post, error)
	FindByID(id int) (*model.Post, error)
	Update(*model.Post) (*model.Post, error)
	Delete(id int) error
	FindAll() ([]*model.Post, error)
}

type postUsecase struct {
	postRepo repository.PostRepository
}

// NewPostUsecase post usecaseのコンストラクタ
func NewPostUsecase(postRepo repository.PostRepository) PostUsecase {
	return &postUsecase{postRepo: postRepo}
}

// Create postを保存するときのユースケース
func (tu *postUsecase) Create(post PostParamsObject) (*model.Post, error) {
	targetPost := model.Post(post)
	err := tu.postRepo.Create(&targetPost)
	if err != nil {
		return nil, err
	}

	return &targetPost, nil
}

// FindByID postをIDで取得するときのユースケース
func (tu *postUsecase) FindByID(id int) (*model.Post, error) {
	foundPost, err := tu.postRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundPost, nil
}

// Update postを更新するときのユースケース
func (tu *postUsecase) Update(post *model.Post) (*model.Post, error) {
	err := tu.postRepo.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// Delete postを削除するときのユースケース
func (tu *postUsecase) Delete(id int) error {
	post, err := tu.postRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.postRepo.Delete(post)
	if err != nil {
		return err
	}

	return nil
}

// FindAll postをIDで取得するときのユースケース
func (tu *postUsecase) FindAll() ([]*model.Post, error) {
	posts, err := tu.postRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
