package persistance

import (
	"api/domain/model"
	"api/domain/repository"

	"gorm.io/gorm"
)

// TagRepository tag repositoryの構造体
type TagRepository struct {
    db *gorm.DB
}

// NewTagRepository tag repositoryのコンストラクタ
func NewTagRepository(db *gorm.DB) repository.TagRepository {
	var tagRepository repository.TagRepository
    tagRepository = &TagRepository{db: db}
	return tagRepository
}

// Create tagの保存
func (tr *TagRepository) Create(tag *model.Tag) error {
    return tr.db.Create(&tag).Error
}

// FindByID tagをIDで取得
func (tr *TagRepository) FindByID(id int) (*model.Tag, error) {
    tag := &model.Tag{ID: id}

    if err := tr.db.First(&tag).Error; err != nil {
        return nil, err
    }

    return tag, nil
}

// Update tagの更新
func (tr *TagRepository) Update(tag *model.Tag) error {
    return tr.db.Model(&model.Tag{ID: tag.ID}).Updates(tag).Error
}

// Delete tagの削除
func (tr *TagRepository) Delete(tag *model.Tag) error {
	return tr.db.Delete(&tag).Error
}

// FindAll tagを全取得
func (tr *TagRepository) FindAll() ([]*model.Tag, error) {
	tags := []*model.Tag{}
	if err := tr.db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
