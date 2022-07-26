package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/src/app/utils"
	"github.com/Sanki0/api-university/src/models"
	"github.com/gorilla/mux"
)


func TestMain(m *testing.M) {
    var application App
    application.Initialize()
	code := m.Run()
	os.Exit(code)
}


func TestCreateStudent(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/student", bytes.NewBuffer([]byte(`{"nombre": "Juan", "Dni": "12345678", "Direccion": "Calle falsa 123", "Fecha_nacimiento": "2020-01-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("INSERT INTO students").
				WithArgs("Juan", "12345678", "Calle falsa 123", "2020-01-01").
				WillReturnResult(sqlmock.NewResult(1, 1))

	app.CreateStudent(w, req)

	if w.Code != 201 {
		t.Fatalf("expected status code to be 201, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

///NO SE SI ESTA CORRECTAMENTE TESTEADO EL CREATE FAIL
func TestCreateStudentFail(t *testing.T) {
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/student", bytes.NewBuffer([]byte(`{"nombre": "Juan", "Dni": "12345678", "Direccion": "Calle falsa 123", "Fecha_nacimiento": "2020-01-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()


	mock.ExpectExec("INSERT INTO students").
				WithArgs("Juan", "12345678", "Calle falsa 123", "2020-01-01").
				WillReturnError(err)

	app.CreateStudent(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

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

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/student", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"nombre", "dni", "direccion","fecha_nacimiento"}).
		AddRow("Juan", "12345678", "Calle falsa 123", "2020-01-01").
        AddRow("Pedro", "87654321", "Calle verdadera 456", "2020-01-02").
        AddRow("Maria", "12345679", "Calle falsa 123", "2020-01-01")


	mock.ExpectQuery("SELECT (.+) FROM students").WillReturnRows(rows)

	 //now we execute our request
	app.GetStudents(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := struct {
		students []models.Student
	}{students: []models.Student{
        {Nombre: "Juan", Dni: "12345678", Direccion: "Calle falsa 123", Fecha_nacimiento: "2020-01-01"},
        {Nombre: "Pedro", Dni: "87654321", Direccion: "Calle verdadera 456", Fecha_nacimiento: "2020-01-02"},
        {Nombre: "Maria", Dni: "12345679", Direccion: "Calle falsa 123", Fecha_nacimiento: "2020-01-01"},
    }}

	utils.AssertJSON(w.Body.Bytes(), data.students, t)

	 //we make sure that all expectations were met
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

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/student", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	mock.ExpectQuery("SELECT (.+) FROM students").WillReturnError(err)

	 //now we execute our request
	app.GetStudents(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	 //we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetStudent(t * testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/student", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"nombre", "dni", "direccion","fecha_nacimiento"}).
		AddRow("Juan", "12345678", "Calle falsa 123", "2020-01-01")

	mock.ExpectQuery("SELECT (.+) FROM students WHERE dni=?").WithArgs("12345678").WillReturnRows(rows)

	vars := map[string]string{
		"dni": "12345678",
	}

	req = mux.SetURLVars(req, vars)

	app.GetStudent(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := models.Student{Nombre: "Juan", Dni: "12345678", Direccion: "Calle falsa 123", Fecha_nacimiento: "2020-01-01"}

	utils.AssertJSON(w.Body.Bytes(), data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateStudent (t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("PUT", "localhost:8080/student", bytes.NewBuffer([]byte(`{"nombre": "Jose", "Dni": "12345678", "Direccion": "Calle falsa 123", "Fecha_nacimiento": "2020-01-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("UPDATE students").
				WithArgs("Jose", "12345678", "Calle falsa 123", "2020-01-01", "12345678").
				WillReturnResult(sqlmock.NewResult(1, 1))
	
	app.UpdateStudent(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

// func TestUpdateStudentFail(t *testing.T){
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	app := &App{DB: db}

// 	req, err := http.NewRequest("PUT", "localhost:8080/student", bytes.NewBuffer([]byte(`{"nombre": "Jose", "Dni": "12345678", "Direccion": "Calle falsa 123", "Fecha_nacimiento": "2020-01-01"}`)))

// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected while creating request", err)
// 	}
// 	w := httptest.NewRecorder()

// 	mock.ExpectExec("UPDATE students").
// 				WithArgs("Jose", "12345678", "Calle falsa 123", "2020-01-01", "12345678").
// 				WillReturnError(err)
	
// 	app.UpdateStudent(w, req)

// 	if w.Code != 404 {
// 		t.Fatalf("expected status code to be 404, but got: %d", w.Code)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// }


func TestDeleteStudent(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("DELETE", "localhost:8080/student/12345678", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	vars := map[string]string{
        "dni": "12345678",
    }

	req = mux.SetURLVars(req, vars)

	mock.ExpectExec("DELETE FROM students").WithArgs("12345678").WillReturnResult(sqlmock.NewResult(1, 1))

	app.DeleteStudent(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


