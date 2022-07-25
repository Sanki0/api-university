package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sanki0/api-university/src/models"
)


func TestMain(m *testing.M) {
    var application App
    application.Initialize()
	code := m.Run()
	os.Exit(code)
}




// func TestGetStudent(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	//create app with mocked db, request and response to test
// 	app := &App{DB: db}
// 	req, err := http.NewRequest("GET", "localhost:8080/student/12345678", nil)
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected while creating request", err)
// 	}
// 	w := httptest.NewRecorder()

// 	row := sqlmock.NewRows([]string{"nombre", "dni", "direccion","fecha_nacimiento"}).
// 						AddRow("Juan", "12345678", "Calle falsa 123", "2020-01-01")
// 	mock.ExpectQuery("SELECT (.+) FROM students WHERE id = ?").WithArgs("12345678").WillReturnRows(row)

// 	app.GetStudent(w, req)

// 	if w.Code != 200 {
// 		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
// 	}

// 	data := models.Student{Nombre: "Juan", Dni: "12345678", Direccion: "Calle falsa 123", Fecha_nacimiento: "2020-01-01"}

// 	app.assertJSON(w.Body.Bytes(), data, t)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

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

	app.assertJSON(w.Body.Bytes(), data.students, t)

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




//  func TestCreateStudent(t *testing.T) {

//      var jsonStr = []byte(`{
//          "Nombre": "cesar",
//          "Dni": "11112222",
//          "Direccion": "Jockey",
//          "Fecha_nacimiento": "2001-09-18"
//        } `)
//      req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(jsonStr))
//      req.Header.Set("Content-Type", "application/json")

//      rr := httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)
	
//  	expected := http.StatusCreated
//  	actual := rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }

//      var m map[string]interface{}
//      json.Unmarshal(rr.Body.Bytes(), &m)

//      if m["nombre"] != "test" {
//          t.Errorf("Expected student name to be 'test'. Got '%v'", m["nombre"])
//      }

//  	if m["dni"] != "12345678" {
//  		t.Errorf("Expected student dni to be '12345678'. Got '%v'", m["dni"])
//  	}
//  	if m["direccion"] != "test" {
//  		t.Errorf("Expected student direccion to be 'test'. Got '%v'", m["direccion"])
//  	}
//  	if m["fecha_nacimiento"] != "test" {
//  		t.Errorf("Expected student fecha_nacimiento to be 'test'. Got '%v'", m["fecha_nacimiento"])
//  	}
//  }


//  func TestGetStudent(t *testing.T) {

//      req, _ := http.NewRequest("GET", "/product", nil)
//  	rr := httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)
	
//  	expected := http.StatusOK
//  	actual := rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }


//  }




//  func TestUpdateProduct(t *testing.T) {

//      req, _ := http.NewRequest("GET", "/product/12345678", nil)
//      rr := httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)

//      var originalStudent map[string]interface{}
//      json.Unmarshal(rr.Body.Bytes(), &originalStudent)

//      var jsonStr = []byte(`{"nombre":"test-updated", "dni":"12345678", "direccion":"test updated", "fecha_nacimiento":"test updated"}`)
//      req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonStr))
//      req.Header.Set("Content-Type", "application/json")

//      rr = httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)

//  	expected := http.StatusOK
//  	actual := rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }
//      var m map[string]interface{}
//      json.Unmarshal(rr.Body.Bytes(), &m)

//      if m["dni"] != originalStudent["dni"] {
//          t.Errorf("Expected the dni to remain the same (%v). Got %v", originalStudent["dni"], m["dni"])
//      }
//      if m["nombre"] == originalStudent["nombre"] {
//          t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalStudent["name"], m["name"], m["name"])
//      }
//      if m["direccion"] == originalStudent["direccion"] {
//  		t.Errorf("Expected the direccion to change from '%v' to '%v'. Got '%v'", originalStudent["direccion"], m["direccion"], m["direccion"])
//  	}
//  	if m["fecha_nacimiento"] == originalStudent["fecha_nacimiento"] {
//  		t.Errorf("Expected the fecha_nacimiento to change from '%v' to '%v'. Got '%v'", originalStudent["fecha_nacimiento"], m["fecha_nacimiento"], m["fecha_nacimiento"])
//  	}

//  }


//  func TestDeleteProduct(t *testing.T) {
//      req, _ := http.NewRequest("GET", "/product/12345678", nil)
//      rr := httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)

//  	expected := http.StatusOK
//  	actual := rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }

//      req, _ = http.NewRequest("DELETE", "/product/12345678", nil)
//      rr = httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)

//  	expected = http.StatusOK
//  	actual = rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }

//      req, _ = http.NewRequest("GET", "/product/12345678", nil)
//  	rr = httptest.NewRecorder()
//      a.Router.ServeHTTP(rr, req)

//  	expected = http.StatusNotFound
//  	actual = rr.Code

//  	if expected != actual {
//          t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//      }
//  }


func (a *App) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}