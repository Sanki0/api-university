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

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		CreateCourse struct {
			Nombre      string
			Descripcion string
			Temas       string
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

func TestCreateCourseFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO courses").
		WithArgs("Devmente", "Introduccion", "Go, AWS").
		WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			CreateCourse struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		createCourse(nombre: "Devmente", descripcion: "Introduccion", temas: "Go, AWS"){
			nombre
			descripcion
			temas
		}
	}`, &resp)

	require.Equal(t, nil, resp.Data.CreateCourse.Nulo)

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

	r := Resolver{DB: db}

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

func TestUpdateCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE courses").
		WithArgs("Devmente", "Introduccion", "Go, AWS", "Devmente").
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		UpdateCourse struct {
			Nombre      string
			Descripcion string
			Temas       string
		}
	}

	c.Post(`
	mutation {
		updateCourse(nombre: "Devmente", descripcion: "Introduccion", temas: "Go, AWS"){
			nombre
			descripcion
			temas
		}
	}`, &resp)

	require.Equal(t, "Devmente", resp.UpdateCourse.Nombre)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateCourseFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE courses").
		WithArgs("Devmente", "Introduccion", "Go, AWS", "Devmente").
		WillReturnError(err)
	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			UpdateCourse struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		updateCourse(nombre: "Devmente", descripcion: "Introduccion", temas: "Go, AWS"){
			nombre
			descripcion
			temas
		}
	}`, &resp)

	require.Equal(t, nil, resp.Data.UpdateCourse.Nulo)

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

	r := Resolver{DB: db}

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

	r := Resolver{DB: db}

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
func TestDeleteCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM courses").WithArgs("Devmente").WillReturnResult(sqlmock.NewResult(1, 1))
	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		DeleteCourse struct {
			Nombre      string
			Descripcion string
			Temas       string
		}
	}

	c.MustPost(`
	mutation {
		deleteCourse(nombre: "Devmente"){
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, "Devmente", resp.DeleteCourse.Nombre)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteCourseFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM courses").WithArgs("Devmente").WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			DeleteCourse struct {
				Nulo error
			}
		}
	}

	c.Post(`
	mutation {
		deleteCourse(nombre: "Devmente"){
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.DeleteCourse.Nulo)

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
	r := Resolver{DB: db}

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
	r := Resolver{DB: db}

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

func TestGetCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"nombre", "descripcion", "temas"}).
		AddRow("Devmente", "Introduccion", "Go, AWS")

	mock.ExpectQuery("SELECT (.+) FROM courses WHERE nombre = ?").
		WithArgs("Devmente").
		WillReturnRows(rows)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetCourse struct {
			Nombre      string
			Descripcion string
			Temas       string
		}
	}

	c.MustPost(`
	query {
		getCourse(nombre: "Devmente"){
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, "Devmente", resp.GetCourse.Nombre)
	require.Equal(t, "Introduccion", resp.GetCourse.Descripcion)
	require.Equal(t, "Go, AWS", resp.GetCourse.Temas)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCourseFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM courses WHERE nombre = ?").
		WithArgs("Devmente").
		WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetCourse struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getCourse(nombre: "Devmente"){
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.GetCourse.Nulo)

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

	r := Resolver{DB: db}

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

	r := Resolver{DB: db}

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

func TestGetCourses(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"nombre", "descripcion", "temas"}).
		AddRow("Devmente", "Introduccion", "Go, AWS").
		AddRow("PHP", "Introduccion", "PHP, MySQL").
		AddRow("Java", "Introduccion", "Java, MySQL")

	mock.ExpectQuery("SELECT (.+) FROM courses").WillReturnRows(rows)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		GetCourses []struct {
			Nombre      string
			Descripcion string
			Temas       string
		}
	}

	c.MustPost(`
	query {
		getCourses{
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, "Devmente", resp.GetCourses[0].Nombre)
	require.Equal(t, "Introduccion", resp.GetCourses[0].Descripcion)
	require.Equal(t, "Go, AWS", resp.GetCourses[0].Temas)
	require.Equal(t, "PHP", resp.GetCourses[1].Nombre)
	require.Equal(t, "Introduccion", resp.GetCourses[1].Descripcion)
	require.Equal(t, "PHP, MySQL", resp.GetCourses[1].Temas)
	require.Equal(t, "Java", resp.GetCourses[2].Nombre)
	require.Equal(t, "Introduccion", resp.GetCourses[2].Descripcion)
	require.Equal(t, "Java, MySQL", resp.GetCourses[2].Temas)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCoursesFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM courses").WillReturnError(err)

	r := Resolver{DB: db}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r})))

	var resp struct {
		Errors struct {
			Message string
			Path    string
		}
		Data struct {
			GetCourses struct {
				Nulo error
			}
		}
	}

	c.Post(`
	query {
		getCourses{
			nombre
			descripcion
			temas
		}
	  }`, &resp)

	require.Equal(t, nil, resp.Data.GetCourses.Nulo)

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

	r := Resolver{DB: db}

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

	r := Resolver{DB: db}

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
