package service

import (
	"github.com/DryundeL/blog-graphql/internal/models"
	"github.com/DryundeL/blog-graphql/internal/repository"
)

type BlogService struct {
	postRepo *repository.PostRepository
}

func NewBlogService(postRepo *repository.PostRepository) *BlogService {
	return &BlogService{postRepo: postRepo}
}

func (s *BlogService) CreatePost(title, content string, userID uint) (*models.Post, error) {
	post := &models.Post{Title: title, Content: content, UserID: userID}
	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *BlogService) GetAllPosts() ([]*models.Post, error) {
	return s.postRepo.FindAll()
}
