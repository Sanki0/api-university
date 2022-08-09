package src

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/graph/generated"

	"github.com/stretchr/testify/require"
)

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
