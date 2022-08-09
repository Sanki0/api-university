package src

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/graph/generated"

	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO students").
		WithArgs("jose", "71231231", "jocke", "18/09/01").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateStudent struct {
			Nombre           string
			Dni              string
			Direccion        string
			Fecha_nacimiento string
		}
	}

	c.MustPost(`
	mutation {
		createStudent(nombre: "jose", dni: "71231231", direccion: "jocke", fecha_nacimiento: "18/09/01"){
			nombre
			dni
			direccion
			fecha_nacimiento
	  	}
	}`, &resp)

	require.Equal(t, "jose", resp.CreateStudent.Nombre)
	require.Equal(t, "71231231", resp.CreateStudent.Dni)
	require.Equal(t, "jocke", resp.CreateStudent.Direccion)
	require.Equal(t, "18/09/01", resp.CreateStudent.Fecha_nacimiento)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestCreateStudentFail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO students").
		WithArgs("jose", "71231231", "jocke", "18/09/01").
		WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			CreateStudent struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		createStudent(nombre: "jose", dni: "71231231", direccion: "jocke", fecha_nacimiento: "18/09/01"){
			nombre
			dni
			direccion
			fecha_nacimiento
	  	}
	}`, &resp)

	require.Equal(t, nil, resp.Data.CreateStudent.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE students").
		WithArgs("Jose", "12345678", "Calle falsa 123", "2020-01-01", "12345678").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		UpdateStudent struct {
			Nombre           string
			Dni              string
			Direccion        string
			Fecha_nacimiento string
		}
	}

	c.Post(`
	mutation {
		updateStudent(nombre: "Jose", dni: "12345678", direccion: "Calle falsa 123", fecha_nacimiento: "2020-01-01"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	}`, &resp)

	require.Equal(t, "12345678", resp.UpdateStudent.Dni)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateStudentFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE students").
		WithArgs("Jose", "12345678", "Calle falsa 123", "2020-01-01", "12345678").
		WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			UpdateStudent struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		updateStudent(nombre: "Jose", dni: "12345678", direccion: "Calle falsa 123", fecha_nacimiento: "2020-01-01"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	}`, &resp)

	require.Equal(t, nil, resp.Data.UpdateStudent.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM students").WithArgs("12345678").WillReturnResult(sqlmock.NewResult(1, 1))

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		DeleteStudent struct {
			Nombre           string
			Dni              string
			Direccion        string
			Fecha_nacimiento string
		}
	}

	c.MustPost(`
	mutation {
		deleteStudent(dni: "12345678"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, "12345678", resp.DeleteStudent.Dni)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteStudentFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM students").WithArgs("12345678").WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			DeleteStudent struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		deleteStudent(dni: "12345678"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.DeleteStudent.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStudent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"nombre", "dni", "direccion", "fecha_nacimiento"}).
		AddRow("Juan", "12345678", "Calle falsa 123", "2020-01-01")

	mock.ExpectQuery("SELECT (.+) FROM students WHERE dni=?").WithArgs("12345678").WillReturnRows(rows)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetStudent struct {
			Nombre           string
			Dni              string
			Direccion        string
			Fecha_nacimiento string
		}
	}

	c.MustPost(`
	query {
		getStudent(dni: "12345678"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, "Juan", resp.GetStudent.Nombre)
	require.Equal(t, "12345678", resp.GetStudent.Dni)
	require.Equal(t, "Calle falsa 123", resp.GetStudent.Direccion)
	require.Equal(t, "2020-01-01", resp.GetStudent.Fecha_nacimiento)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStudentFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM students WHERE dni=?").WithArgs("12345678").WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetStudent struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getStudent(dni: "12345678"){
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.GetStudent.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStudents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"nombre", "dni", "direccion", "fecha_nacimiento"}).
		AddRow("Juan", "12345678", "Calle falsa 123", "2020-01-01").
		AddRow("Pedro", "87654321", "Calle verdadera 456", "2020-01-02").
		AddRow("Maria", "12345679", "Calle falsa 123", "2020-01-01")

	mock.ExpectQuery("SELECT (.+) FROM students").WillReturnRows(rows)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetStudents []struct {
			Nombre           string
			Dni              string
			Direccion        string
			Fecha_nacimiento string
		}
	}

	c.MustPost(`
	query {
		getStudents{
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, "Juan", resp.GetStudents[0].Nombre)
	require.Equal(t, "12345678", resp.GetStudents[0].Dni)
	require.Equal(t, "Calle falsa 123", resp.GetStudents[0].Direccion)
	require.Equal(t, "2020-01-01", resp.GetStudents[0].Fecha_nacimiento)
	require.Equal(t, "Pedro", resp.GetStudents[1].Nombre)
	require.Equal(t, "87654321", resp.GetStudents[1].Dni)
	require.Equal(t, "Calle verdadera 456", resp.GetStudents[1].Direccion)
	require.Equal(t, "2020-01-02", resp.GetStudents[1].Fecha_nacimiento)
	require.Equal(t, "Maria", resp.GetStudents[2].Nombre)
	require.Equal(t, "12345679", resp.GetStudents[2].Dni)
	require.Equal(t, "Calle falsa 123", resp.GetStudents[2].Direccion)
	require.Equal(t, "2020-01-01", resp.GetStudents[2].Fecha_nacimiento)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStudentsFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM students").WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetStudents struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getStudents{
			nombre
			dni
			direccion
			fecha_nacimiento
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.GetStudents.Nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
