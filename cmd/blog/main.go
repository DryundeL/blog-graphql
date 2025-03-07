package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DryundeL/blog-graphql/internal/config"
	"github.com/DryundeL/blog-graphql/internal/db"
	"github.com/DryundeL/blog-graphql/internal/graphql/generated"
	"github.com/DryundeL/blog-graphql/internal/graphql/resolver"
	"github.com/DryundeL/blog-graphql/internal/middleware"
	"github.com/DryundeL/blog-graphql/internal/repository"
	"github.com/DryundeL/blog-graphql/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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

	// Настройка HTTP-сервера с Gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	gqlHandler := handler.New(
		generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}),
	)
	r.POST("/query", middleware.AuthMiddleware(cfg), func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Запуск сервера
	log.Printf("Server running on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
