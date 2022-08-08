package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sanki0/api-university/graph/model"
)

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

// DeleteRecord is the resolver for the deleteRecord field.
func (r *mutationResolver) DeleteRecord(ctx context.Context, student string, course string) (*model.Record, error) {

	record := model.Record{Student: student, Course: course}

	_, err := r.DB.Exec("DELETE FROM records WHERE student = ? AND course = ?", student, course)

	if err != nil {
		return nil, err
	}

	return &record, nil
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
