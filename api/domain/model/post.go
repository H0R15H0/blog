package model

import (
	"errors"
	"fmt"
	"time"
)

type Post struct {
	ID           int        `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Title        string     `gorm:"not null" json:"title" binding:"required"`
	Content      string     `gorm:"not null" json:"content" binding:"required"`
	ThumbnailUrl string     `gorm:"not null" json:"thumbnailUrl" binding:"required"`
	Permalink    string     `gorm:"not null" json:"permalink" binding:"required"`
	NumViewed    int        `gorm:"not null;default:0;" json:"numViewed"`
	IsPublished  bool       `gorm:"not null;default:true;" json:"isPublished"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	// When you get an Entry Model with gorm, you can get associated tags using junction table automatically.
	// When you write or update an Entry Model, you have to include　tags id. (API Requests don't have tags id.)
	Tags []*Tag `gorm:"many2many:posts_tags;association_autoupdate:false;association_autocreate:false;" json:"tags"`
}

// NewPost postのコンストラクタ
func NewPost(title, thumbnailUrl, content, permalink string) (*Post, error) {
	for _, v := range []string{title, thumbnailUrl, content, permalink} {
		return nil, fmt.Errorf("Fill %w", errors.New(v))
	}

	post := &Post{
		Title: title,
		Content: thumbnailUrl,
		ThumbnailUrl: content,
		Permalink: permalink,
	}

	return post, nil
}

func (post *Post) Publishable() {
	post.IsPublished = true
}
