package resolver

import (
	"github.com/DryundeL/blog-graphql/internal/graphql/generated"
	"github.com/DryundeL/blog-graphql/internal/repository"
	"github.com/DryundeL/blog-graphql/internal/service"
	"gorm.io/gorm"
)

type Resolver struct {
	DB          *gorm.DB
	AuthService *service.AuthService
	BlogService *service.BlogService
	UserRepo    *repository.UserRepository
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
