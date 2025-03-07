package resolver

import (
	"context"
	"fmt"
	"github.com/DryundeL/blog-graphql/internal/models"
)

func (r *mutationResolver) Register(ctx context.Context, username, password, email string) (*models.User, error) {
	return r.AuthService.Register(username, password, email)
}

func (r *mutationResolver) Login(ctx context.Context, username, password string) (string, error) {
	return r.AuthService.Login(username, password)
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	userID, _ := ctx.Value("userID").(uint)
	return r.UserRepo.FindByID(userID)
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil // Преобразуем uint в string
}
