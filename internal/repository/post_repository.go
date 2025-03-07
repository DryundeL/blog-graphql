package repository

import (
	"github.com/DryundeL/blog-graphql/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) FindAll() ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Preload("User").Find(&posts).Error
	return posts, err
}
