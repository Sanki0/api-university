package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/src/app/utils"
	"github.com/Sanki0/api-university/src/models"
	"github.com/gorilla/mux"
)


func TestCreateCourse(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/course", 
		bytes.NewBuffer([]byte(`{"Nombre": "Devmente","Descripcion":"Introduccion","Temas":"Go, AWS"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("INSERT INTO courses").
				WithArgs("Devmente", "Introduccion", "Go, AWS").
				WillReturnResult(sqlmock.NewResult(1, 1))

	app.CreateCourse(w, req)

	if w.Code != 201 {
		t.Fatalf("expected status code to be 201, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestCreateCourseFailPayload(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/course", 
		bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()


	app.CreateCourse(w, req)

	if w.Code != 400 {
		t.Fatalf("expected status code to be 400, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

///NO SE SI ESTA CORRECTAMENTE TESTEADO EL CREATE FAIL
func TestCreateCourseFailServerError(t *testing.T) {
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/course", 
					bytes.NewBuffer([]byte(`{"Nombre": "Devmente","Descripcion":"Introduccion","Temas":"Go, AWS"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()


	mock.ExpectExec("INSERT INTO courses").
				WithArgs("Devmente", "Introduccion", "Go, AWS").
				WillReturnError(err)

	app.CreateCourse(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

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

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/course", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"nombre", "descripcion", "temas"}).
		AddRow("Devmente", "Introduccion", "Go, AWS").
		AddRow("PHP", "Introduccion", "PHP, MySQL").
		AddRow("Java", "Introduccion", "Java, MySQL")



	mock.ExpectQuery("SELECT (.+) FROM courses").WillReturnRows(rows)

	 //now we execute our request
	app.GetCourses(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := []models.Course{
        { Nombre: "Devmente", Descripcion: "Introduccion", Temas: "Go, AWS"},
		{ Nombre: "PHP", Descripcion: "Introduccion", Temas: "PHP, MySQL"},
		{ Nombre: "Java", Descripcion: "Introduccion", Temas: "Java, MySQL"},
    }

	utils.AssertJSON(w.Body.Bytes(), data, t)

	 //we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


func TestGetCoursesFailServerError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/course", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	mock.ExpectQuery("SELECT (.+) FROM courses").WillReturnError(err)

	 //now we execute our request
	app.GetCourses(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	 //we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetCourse(t * testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/course", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"nombre", "descripcion", "temas"}).
		AddRow("Devmente", "Introduccion", "Go, AWS")

	mock.ExpectQuery("SELECT (.+) FROM courses WHERE nombre = ?").
		WithArgs("Devmente").
		WillReturnRows(rows)

	vars := map[string]string{
		"nombre": "Devmente",
	}

	req = mux.SetURLVars(req, vars)

	app.GetCourse(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := models.Course{ Nombre: "Devmente", Descripcion: "Introduccion", Temas: "Go, AWS"}

	utils.AssertJSON(w.Body.Bytes(), data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCourseFailNotFound(t * testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/course", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"nombre", "descripcion", "temas"})

	mock.ExpectQuery("SELECT (.+) FROM courses WHERE nombre = ?").
		WithArgs("Devmente").
		WillReturnRows(rows)

	vars := map[string]string{
		"nombre": "Devmente",
	}

	req = mux.SetURLVars(req, vars)

	app.GetCourse(w, req)

	if w.Code != 404 {
		t.Fatalf("expected status code to be 404, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCourseFailServerError(t * testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/course", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()


	mock.ExpectQuery("SELECT (.+) FROM courses WHERE nombre = ?").
		WithArgs("Devmente").
		WillReturnError(err)

	vars := map[string]string{
		"nombre": "Devmente",
	}

	req = mux.SetURLVars(req, vars)

	app.GetCourse(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateCourse (t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("PUT", "localhost:8080/course", 
					bytes.NewBuffer([]byte(`{"Nombre": "Devmente","Descripcion":"Introduccion","Temas":"Go, AWS"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("UPDATE courses").
				WithArgs("Devmente", "Introduccion", "Go, AWS","Devmente").
				WillReturnResult(sqlmock.NewResult(1, 1))
	
	app.UpdateCourse(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateCourseFailPayload (t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("PUT", "localhost:8080/course", 
					bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	
	app.UpdateCourse(w, req)

	if w.Code != 400 {
		t.Fatalf("expected status code to be 400, but got: %d", w.Code)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateCourseFailServerError (t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("PUT", "localhost:8080/course", 
					bytes.NewBuffer([]byte(`{"Nombre": "Devmente","Descripcion":"Introduccion","Temas":"Go, AWS"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("UPDATE courses").
				WithArgs("Devmente", "Introduccion", "Go, AWS","Devmente").
				WillReturnError(err)
	
	app.UpdateCourse(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}


func TestDeleteCourse(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("DELETE", "localhost:8080/course/Devmente", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	vars := map[string]string{
        "nombre": "Devmente",
    }

	req = mux.SetURLVars(req, vars)

	mock.ExpectExec("DELETE FROM courses").WithArgs("Devmente").WillReturnResult(sqlmock.NewResult(1, 1))

	app.DeleteCourse(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteCourseFailServerError(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("DELETE", "localhost:8080/course/Devmente", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	vars := map[string]string{
        "nombre": "Devmente",
    }

	req = mux.SetURLVars(req, vars)

	mock.ExpectExec("DELETE FROM courses").WithArgs("Devmente").WillReturnError(err)

	app.DeleteCourse(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}