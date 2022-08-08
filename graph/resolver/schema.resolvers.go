package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sanki0/api-university/graph/model"
)

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {

	course := model.Course{Nombre: nombre, Descripcion: descripcion, Temas: temas}

	_, err := r.DB.Exec("INSERT INTO courses(nombre, descripcion, temas) VALUES(?,?,?)",
		nombre, descripcion, temas)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

// CreateRecord is the resolver for the createRecord field.
func (r *mutationResolver) CreateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course, Startdate: startdate, Finishdate: finishdate}

	_, err := r.DB.Exec("INSERT INTO records(student, course, startdate, finishdate) VALUES(?, ?, ?, ?)",
		student, course, startdate, finishdate)

	if err != nil {
		return nil, err
	}
	return &record, nil
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {

	course := model.Course{Nombre: nombre}

	_, err :=
		r.DB.Exec("UPDATE courses SET nombre = ?, descripcion = ?, temas = ? WHERE nombre = ?",
			nombre, descripcion, temas, nombre)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

// UpdateRecord is the resolver for the updateRecord field.
func (r *mutationResolver) UpdateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	_, err :=
		r.DB.Exec("UPDATE records SET student = ?, course = ?, startdate = ?, finishdate = ? WHERE student = ? AND course = ?",
			student, course, startdate, finishdate, student, course)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, nombre string) (*model.Course, error) {

	course := model.Course{Nombre: nombre}

	_, err := r.DB.Exec("DELETE FROM courses WHERE nombre = ?", nombre)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

// DeleteRecord is the resolver for the deleteRecord field.
func (r *mutationResolver) DeleteRecord(ctx context.Context, student string, course string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	_, err := r.DB.Exec("DELETE FROM records WHERE student = ? AND course = ?", student, course)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {

	var courses []*model.Course
	rows, err := r.DB.Query("SELECT * FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c model.Course
		if err := rows.Scan(&c.Nombre, &c.Descripcion, &c.Temas); err != nil {
			return nil, err
		}
		courses = append(courses, &c)
	}

	return courses, nil
}

// GetRecords is the resolver for the getRecords field.
func (r *queryResolver) GetRecords(ctx context.Context) ([]*model.Record, error) {

	var records []*model.Record
	rows, err := r.DB.Query("SELECT * FROM records")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var r model.Record
		if err := rows.Scan(&r.Student, &r.Course, &r.Startdate, &r.Finishdate); err != nil {
			return nil, err
		}
		records = append(records, &r)
	}

	return records, nil
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, nombre string) (*model.Course, error) {

	course := model.Course{Nombre: nombre}

	err := r.DB.QueryRow("SELECT * FROM courses WHERE nombre = ?",
		nombre).Scan(&course.Nombre, &course.Descripcion, &course.Temas)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

// GetRecord is the resolver for the getRecord field.
func (r *queryResolver) GetRecord(ctx context.Context, student string, course string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	err := r.DB.QueryRow("SELECT * FROM records WHERE student = ? AND course = ?",
		student, course).Scan(&record.Student, &record.Course, &record.Startdate, &record.Finishdate)

	if err != nil {
		return nil, err
	}

	return &record, nil
}
