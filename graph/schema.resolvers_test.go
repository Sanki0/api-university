package graph

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

	c:= client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateStudent struct {
			Nombre string
			Dni string
			Direccion string
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

func TestCreateCourse(t *testing.T) {
	t.Skip("not implemented")
}

func TestCreateRecord(t *testing.T) {
	t.Skip("not implemented")
}

func TestUpdateStudent(t *testing.T) {
	t.Skip("not implemented")
}

func TestUpdateCourse(t *testing.T) {
	t.Skip("not implemented")
}

func TestUpdateRecord(t *testing.T) {
	t.Skip("not implemented")
}

func TestDeleteStudent(t *testing.T) {
	t.Skip("not implemented")
}

func TestDeleteCourse(t *testing.T) {
	t.Skip("not implemented")
}

func TestDeleteRecord(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetStudent(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetCourse(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetRecord(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetStudents(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetCourses(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetRecords(t *testing.T) {
	t.Skip("not implemented")
}


