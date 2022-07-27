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


func TestCreateRecord(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/record", 
		bytes.NewBuffer([]byte(`{"Student": "12453124","Course": "PHP","Startdate": "2022-09-01","Finishdate": "2022-12-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("INSERT INTO records").
				WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01").
				WillReturnResult(sqlmock.NewResult(1, 1))

	app.CreateRecord(w, req)

	if w.Code != 201 {
		t.Fatalf("expected status code to be 201, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

///NO SE SI ESTA CORRECTAMENTE TESTEADO EL CREATE FAIL
func TestCreateRecordFail(t *testing.T) {
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("POST", "localhost:8080/record", 
					bytes.NewBuffer([]byte(`{"Student": "12453124","Course": "PHP","Startdate": "2022-09-01","Finishdate": "2022-12-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()


	mock.ExpectExec("INSERT INTO records").
				WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01").
				WillReturnError(err)

	app.CreateRecord(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

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

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/record", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"student", "course", "startdate", "finishdate"}).
		AddRow("12453124", "PHP", "2022-09-01", "2022-12-01").
		AddRow("12453124", "Java", "2022-09-01", "2022-12-01").
		AddRow("12453124", "Python", "2022-09-01", "2022-12-01")



	mock.ExpectQuery("SELECT (.+) FROM records").WillReturnRows(rows)

	 //now we execute our request
	app.GetRecords(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := []models.Record{
        {Student: "12453124", Course: "PHP", Startdate: "2022-09-01", Finishdate: "2022-12-01"},
		{Student: "12453124", Course: "Java", Startdate: "2022-09-01", Finishdate: "2022-12-01"},
		{Student: "12453124", Course: "Python", Startdate: "2022-09-01", Finishdate: "2022-12-01"},
    }

	utils.AssertJSON(w.Body.Bytes(), data, t)

	 //we make sure that all expectations were met
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

	 //create app with mocked db, request and response to test
	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/record", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	mock.ExpectQuery("SELECT (.+) FROM records").WillReturnError(err)

	 //now we execute our request
	app.GetRecords(w, req)

	if w.Code != 500 {
		t.Fatalf("expected status code to be 500, but got: %d", w.Code)
	}

	 //we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetRecord(t * testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("GET", "localhost:8080/record", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"student", "course", "startdate", "finishdate"}).
		AddRow("12453124", "PHP", "2022-09-01", "2022-12-01")

	mock.ExpectQuery ("SELECT (.+) FROM records WHERE student = ?").
		WithArgs("12453124", "PHP").
		WillReturnRows(rows)

	vars := map[string]string{
		"student": "12453124",
		"course": "PHP",
	}

	req = mux.SetURLVars(req, vars)

	app.GetRecord(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := models.Record{ Student: "12453124", Course: "PHP", Startdate: "2022-09-01", Finishdate: "2022-12-01"}

	utils.AssertJSON(w.Body.Bytes(), data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateRecord (t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("PUT", "localhost:8080/record", 
					bytes.NewBuffer([]byte(`{"student": "12453124", "course": "PHP", "startdate": "2022-09-01", "finishdate": "2022-12-01"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	mock.ExpectExec("UPDATE records").
				WithArgs("12453124", "PHP", "2022-09-01", "2022-12-01", "12453124", "PHP").
				WillReturnResult(sqlmock.NewResult(1, 1))
	
	app.UpdateRecord(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}


func TestDeleteRecord(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := &App{DB: db}
	req, err := http.NewRequest("DELETE", "localhost:8080/record/12453124/PHP", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	vars := map[string]string{
        "student": "12453124",
		"course": "PHP",
    }

	req = mux.SetURLVars(req, vars)

	mock.ExpectExec("DELETE FROM records").WithArgs("12453124","PHP").WillReturnResult(sqlmock.NewResult(1, 1))

	app.DeleteRecord(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}