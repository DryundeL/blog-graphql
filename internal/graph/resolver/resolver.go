package resolver

import (
	"github.com/DryundeL/blog-graphql/internal/repository"
	"github.com/DryundeL/blog-graphql/internal/service"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          *gorm.DB
	AuthService *service.AuthService
	BlogService *service.BlogService
	UserRepo    *repository.UserRepository
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
