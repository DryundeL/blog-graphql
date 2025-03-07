package resolver

import (
	"context"
	"github.com/DryundeL/blog-graphql/internal/models"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title, content string) (*models.Post, error) {
	userID, _ := ctx.Value("userID").(uint)
	return r.BlogService.CreatePost(title, content, userID)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.BlogService.GetAllPosts()
}
