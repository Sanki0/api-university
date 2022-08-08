package test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/graph/generated"
	"github.com/Sanki0/api-university/graph/resolver"

	"github.com/stretchr/testify/require"
)

func TestCreateRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO records").
		WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateRecord struct {
			Student    string
			Course     string
			Startdate  string
			Finishdate string
		}
	}

	c.MustPost(`
	mutation {
		createRecord(student: "12453124", course: "PHP", startdate: "2022-09-01", finishdate: "2022-12-01"){
			student
			course
			startdate
			finishdate
		}
	}`, &resp)

	require.Equal(t, "12453124", resp.CreateRecord.Student)
	require.Equal(t, "PHP", resp.CreateRecord.Course)
	require.Equal(t, "2022-09-01", resp.CreateRecord.Startdate)
	require.Equal(t, "2022-12-01", resp.CreateRecord.Finishdate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestCreateRecordFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO records").
		WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01").
		WillReturnError(err)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			CreateRecord struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		createRecord(student: "12453124", course: "PHP", startdate: "2022-09-01", finishdate: "2022-12-01"){
			student
			course
			startdate
			finishdate
		}
	}`, &resp)

	require.Equal(t, nil, resp.Data.CreateRecord.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE records").
		WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01", "12453124", "PHP").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		UpdateRecord struct {
			Student    string
			Course     string
			Startdate  string
			Finishdate string
		}
	}

	c.Post(`
	mutation {
		updateRecord(student: "12453124", course: "PHP", startdate: "2022-09-01", finishdate: "2022-12-01"){
		  student
		  course
		  startdate
		  finishdate
		}
	  }`, &resp)

	require.Equal(t, "12453124", resp.UpdateRecord.Student)
	require.Equal(t, "PHP", resp.UpdateRecord.Course)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateRecordFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE records").
		WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01", "12453124", "PHP").
		WillReturnError(err)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			UpdateRecord struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		updateRecord(student: "12453124", course: "PHP", startdate: "2022-09-01", finishdate: "2022-12-01"){
		  student
		  course
		  startdate
		  finishdate
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.UpdateRecord.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM records").WithArgs("12453124", "PHP").WillReturnResult(sqlmock.NewResult(1, 1))
	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		DeleteRecord struct {
			Student    string
			Course     string
			Startdate  string
			Finishdate string
		}
	}

	c.MustPost(`
	mutation {
		deleteRecord(student: "12453124", course: "PHP"){
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	require.Equal(t, "12453124", resp.DeleteRecord.Student)
	require.Equal(t, "PHP", resp.DeleteRecord.Course)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteRecordFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM records").WithArgs("12453124", "PHP").WillReturnError(err)
	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			DeleteRecord struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		deleteRecord(student: "12453124", course: "PHP"){
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.DeleteRecord.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"student", "course", "startdate", "finishdate"}).
		AddRow("12453124", "PHP", "2022-09-01", "2022-12-01")

	mock.ExpectQuery("SELECT (.+) FROM records WHERE student = ?").
		WithArgs("12453124", "PHP").
		WillReturnRows(rows)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetRecord struct {
			Student    string
			Course     string
			Startdate  string
			Finishdate string
		}
	}

	c.MustPost(`
	query {
		getRecord(student: "12453124", course: "PHP"){
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	require.Equal(t, "12453124", resp.GetRecord.Student)
	require.Equal(t, "PHP", resp.GetRecord.Course)
	require.Equal(t, "2022-09-01", resp.GetRecord.Startdate)
	require.Equal(t, "2022-12-01", resp.GetRecord.Finishdate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRecordFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM records WHERE student = ?").
		WithArgs("12453124", "PHP").
		WillReturnError(err)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetRecord struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getRecord(student: "12453124", course: "PHP"){
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRecords(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"student", "course", "startdate", "finishdate"}).
		AddRow("12453124", "PHP", "2022-09-01", "2022-12-01").
		AddRow("12453124", "Java", "2022-09-01", "2022-12-01").
		AddRow("12453124", "Python", "2022-09-01", "2022-12-01")

	mock.ExpectQuery("SELECT (.+) FROM records").WillReturnRows(rows)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetRecords []struct {
			Student    string
			Course     string
			Startdate  string
			Finishdate string
		}
	}

	c.MustPost(`
	query {
		getRecords{
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	require.Equal(t, "12453124", resp.GetRecords[0].Student)
	require.Equal(t, "PHP", resp.GetRecords[0].Course)
	require.Equal(t, "2022-09-01", resp.GetRecords[0].Startdate)
	require.Equal(t, "2022-12-01", resp.GetRecords[0].Finishdate)
	require.Equal(t, "12453124", resp.GetRecords[1].Student)
	require.Equal(t, "Java", resp.GetRecords[1].Course)
	require.Equal(t, "2022-09-01", resp.GetRecords[1].Startdate)
	require.Equal(t, "2022-12-01", resp.GetRecords[1].Finishdate)
	require.Equal(t, "12453124", resp.GetRecords[2].Student)
	require.Equal(t, "Python", resp.GetRecords[2].Course)
	require.Equal(t, "2022-09-01", resp.GetRecords[2].Startdate)
	require.Equal(t, "2022-12-01", resp.GetRecords[2].Finishdate)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRecordsFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM records").WillReturnError(err)

	r := resolver.Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetRecords struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getRecords{
			student
			course
			startdate
			finishdate
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.GetRecords.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
