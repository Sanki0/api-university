package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sanki0/api-university/graph/generated"

	"github.com/Sanki0/api-university/graph/model"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	student := model.Student{Nombre: nombre, Dni: dni, Direccion: direccion, FechaNacimiento: fechaNacimiento}

	r.DB.Create(&student)

	return &student, nil

}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {


	course := model.Course{Nombre: nombre, Descripcion: descripcion, Temas: temas}

	r.DB.Create(&course)

	return &course, nil
}

// CreateRecord is the resolver for the createRecord field.
func (r *mutationResolver) CreateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {


	record := model.Record{Student: student, Course: course, Startdate: startdate, Finishdate: finishdate}

	r.DB.Create(&record)

	return &record, nil
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {

	student := model.Student{Dni: dni}

	r.DB.Model(&student).Where("dni = ?", dni).Update(&model.Student{Nombre: nombre, Direccion: direccion, FechaNacimiento: fechaNacimiento})

	return &student, nil

}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {


	course := model.Course{Nombre: nombre}

	r.DB.Model(&course).Where("nombre = ?", nombre).Update(&model.Course{Descripcion: descripcion, Temas: temas})

	return &course, nil
}

// UpdateRecord is the resolver for the updateRecord field.
func (r *mutationResolver) UpdateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {


	record := model.Record{Student: student, Course: course}

	r.DB.Model(&record).Where("student = ? AND course = ?", student, course).Update(&model.Record{Startdate: startdate, Finishdate: finishdate})

	return &record, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, dni string) (*model.Student, error) {


	student := model.Student{Dni: dni}

	r.DB.Where("dni = ?", dni).First(&student).Delete(&student)

	return &student, nil
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, nombre string) (*model.Course, error) {

	course := model.Course{Nombre: nombre}

	r.DB.Where("nombre = ?", nombre).First(&course).Delete(&course)

	return &course, nil
}

// DeleteRecord is the resolver for the deleteRecord field.
func (r *mutationResolver) DeleteRecord(ctx context.Context, student string, course string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	r.DB.Where("student = ? AND course = ?", student, course).First(&record).Delete(&record)

	return &record, nil
}

// GetStudents is the resolver for the getStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Student, error) {

	var students []*model.Student
	r.DB.Find(&students)

	return students, nil
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {

	var courses []*model.Course
	r.DB.Find(&courses)

	return courses, nil
}

// GetRecords is the resolver for the getRecords field.
func (r *queryResolver) GetRecords(ctx context.Context) ([]*model.Record, error) {

	var records []*model.Record
	r.DB.Find(&records)

	return records, nil
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, dni string) (*model.Student, error) {

	student := model.Student{Dni: dni}

	r.DB.Where("dni = ?", dni).First(&student)

	return &student, nil
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, nombre string) (*model.Course, error) {

	course := model.Course{Nombre: nombre}

	r.DB.Where("nombre = ?", nombre).First(&course)

	return &course, nil
}

// GetRecord is the resolver for the getRecord field.
func (r *queryResolver) GetRecord(ctx context.Context, student string, course string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	r.DB.Where("student = ? AND course = ?", student, course).First(&record)

	return &record, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }



/*

class mutationResolver{
	private:
		Resolver *resolver;
	public:
		mutationResolver(Resolver *resolver){
			this->resolver = resolver;
		create
		create
		create
		update
		delete

}

class queryResolver{
	private:
		Resolver *resolver;
	public:
		queryResolver(Resolver *resolver){
			this->resolver = resolver;
		get
		get
		get
		get
		get
		get
		get






*/
