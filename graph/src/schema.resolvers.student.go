package src

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sanki0/api-university/graph/model"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	student := model.Student{Nombre: nombre, Dni: dni, Direccion: direccion, FechaNacimiento: fechaNacimiento}

	_, err := r.DB.Exec("INSERT INTO students(Nombre, Dni, Direccion, Fecha_nacimiento) VALUES(?, ?, ?, ?)",
		student.Nombre, student.Dni, student.Direccion, student.FechaNacimiento)

	if err != nil {
		return nil, err
	}

	return &student, nil

}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {

	student := model.Student{Dni: dni}

	_, err := r.DB.Exec("UPDATE students SET Nombre=?, Dni=?, Direccion=?, Fecha_nacimiento=? WHERE dni=?",
		nombre, dni, direccion, fechaNacimiento, dni)

	if err != nil {
		return nil, err
	}

	return &student, nil

}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, dni string) (*model.Student, error) {

	student := model.Student{Dni: dni}

	_, err := r.DB.Exec("DELETE FROM students WHERE dni = ?", dni)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

// GetStudents is the resolver for the getStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Student, error) {

	var students []*model.Student
	rows, err := r.DB.Query("SELECT * FROM students")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.FechaNacimiento); err != nil {
			return nil, err
		}
		students = append(students, &s)
	}

	return students, nil
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, dni string) (*model.Student, error) {

	student := model.Student{Dni: dni}

	err := r.DB.QueryRow("SELECT * FROM students WHERE dni = ?",
		dni).Scan(&student.Nombre, &student.Dni, &student.Direccion, &student.FechaNacimiento)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
