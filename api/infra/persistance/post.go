package persistance

import (
	"api/domain/model"
	"api/domain/repository"

	"gorm.io/gorm"
)

// PostRepository post repositoryの構造体
type PostRepository struct {
    db *gorm.DB
}

// NewPostRepository post repositoryのコンストラクタ
func NewPostRepository(db *gorm.DB) repository.PostRepository {
	var postRepository repository.PostRepository
    postRepository = &PostRepository{db: db}
	return postRepository
}

// Create postの保存
func (tr *PostRepository) Create(post *model.Post) error {
    return tr.db.Create(&post).Error
}

// FindByID postをIDで取得
func (tr *PostRepository) FindByID(id int) (*model.Post, error) {
    post := &model.Post{ID: id}

    if err := tr.db.First(&post).Error; err != nil {
        return nil, err
    }

    return post, nil
}

// Update postの更新
func (tr *PostRepository) Update(post *model.Post) error {
    return tr.db.Model(&model.Post{ID: post.ID}).Updates(post).Error
}

// Delete postの削除
func (tr *PostRepository) Delete(post *model.Post) error {
	return tr.db.Delete(&post).Error
}

// FindAll postを全取得
func (tr *PostRepository) FindAll() ([]*model.Post, error) {
	posts := []*model.Post{}
	if err := tr.db.Preload("Tags").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
