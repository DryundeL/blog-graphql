package resolver

import (
	"context"
	"fmt"
	"github.com/DryundeL/blog-graphql/internal/models"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title, content string) (*models.Post, error) {
	userID, _ := ctx.Value("userID").(uint)
	return r.BlogService.CreatePost(title, content, userID)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.BlogService.GetAllPosts()
}

func (r *postResolver) Author(ctx context.Context, obj *models.Post) (*models.User, error) {
	user, err := r.UserRepo.FindByID(obj.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil // Преобразуем uint в string
}
