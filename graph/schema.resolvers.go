package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Sanki0/api-university/graph/generated"
	"github.com/Sanki0/api-university/graph/model"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateRecord is the resolver for the createRecord field.
func (r *mutationResolver) CreateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateRecord is the resolver for the updateRecord field.
func (r *mutationResolver) UpdateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, dni string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, nombre string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteRecord is the resolver for the deleteRecord field.
func (r *mutationResolver) DeleteRecord(ctx context.Context, student string, course string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetStudents is the resolver for the getStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetRecords is the resolver for the getRecords field.
func (r *queryResolver) GetRecords(ctx context.Context) ([]*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, dni string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, nombre string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetRecord is the resolver for the getRecord field.
func (r *queryResolver) GetRecord(ctx context.Context, student string, course string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
