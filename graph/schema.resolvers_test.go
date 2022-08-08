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

	c:= client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct{
			Message string
			Path string
		}
		Data struct{
			CreateStudent struct{
				nulo error
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
	
	require.Equal(t, nil, resp.Data.CreateStudent.nulo)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}



func TestCreateCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	

	mock.ExpectExec("INSERT INTO courses").
	WithArgs("Devmente", "Introduccion", "Go, AWS").
	WillReturnResult(sqlmock.NewResult(1, 1))
	
	r := Resolver{DB: db}

	c:= client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateCourse struct {
			Nombre string
			Descripcion string
			Temas string
		}
	}

	c.MustPost(`
	mutation {
		createCourse(nombre: "Devmente", descripcion: "Introduccion", temas: "Go, AWS"){
			nombre
			descripcion
			temas
		}
	}`, &resp)
	


	require.Equal(t, "Devmente", resp.CreateCourse.Nombre)
	require.Equal(t, "Introduccion", resp.CreateCourse.Descripcion)
	require.Equal(t, "Go, AWS", resp.CreateCourse.Temas)


	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	
	mock.ExpectExec("INSERT INTO records").
				WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01").
				WillReturnResult(sqlmock.NewResult(1, 1))
	
	r := Resolver{DB: db}

	c:= client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateRecord struct {
			Student string
			Course string
			Startdate string
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


