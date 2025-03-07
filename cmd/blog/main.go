package main

import (
	"github.com/DryundeL/blog-graphql/internal/config"
	"github.com/DryundeL/blog-graphql/internal/db"
	"github.com/DryundeL/blog-graphql/internal/graphql/resolver"
	"github.com/DryundeL/blog-graphql/internal/repository"
	"github.com/DryundeL/blog-graphql/internal/service"
)

func main() {
	//Init config
	cfg := config.LoadConfig()

	// Run migrations
	db.RunMigrations(cfg)

	// Init db
	database := db.InitDB(cfg)

	// Define repositories
	userRepo := repository.NewUserRepository(database)
	postRepo := repository.NewPostRepository(database)

	// Сервисы
	authSvc := service.NewAuthService(database, cfg)
	blogSvc := service.NewBlogService(postRepo)

	// GraphQL резолверы
	resolvers := &resolver.Resolver{
		AuthService: authSvc,
		BlogService: blogSvc,
		UserRepo:    userRepo,
	}
}
