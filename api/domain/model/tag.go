package model

import (
	"time"
)

type Tag struct {
	ID        int        `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Name      string     `gorm:"unique;not null" json:"name" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`

	Posts []*Post `gorm:"many2many:posts_tags;save_associations:false;association_autocreate:false;" json:"posts"`
}
