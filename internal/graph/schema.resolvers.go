package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"github.com/DryundeL/blog-graphql/internal/graph/resolver"

	"github.com/DryundeL/blog-graphql/internal/graph/model"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - CreatePost"))
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, postID int32, input *model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - UpdatePost"))
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetAllPosts - GetAllPosts"))
}

// GetOnePost is the resolver for the GetOnePost field.
func (r *queryResolver) GetOnePost(ctx context.Context, id int32) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetOnePost - GetOnePost"))
}

// Mutation returns MutationResolver implementation.
func (r *resolver.Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *resolver.Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *resolver.Resolver }
type queryResolver struct{ *resolver.Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
*/
