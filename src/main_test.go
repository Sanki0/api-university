package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Sanki0/api-university/src/app"
)

var a app.App

func TestMain(m *testing.M) {
    a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetNonExistentStudent(t *testing.T) {

    req, _ := http.NewRequest("GET", "/student/11112222", nil)
	rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
	
	expected := http.StatusNotFound
	actual := rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }

    var m map[string]string
    json.Unmarshal(rr.Body.Bytes(), &m)
	
    if m["error"] != "Student not found" {
        t.Errorf("Expected the 'error' key of the response to be set to 'Student not found'. Got '%s'", m["error"])
    }
}

func TestCreateStudent(t *testing.T) {

    var jsonStr = []byte(`{"nombre":"test", "dni":"12345678", "direccion":"test", "fecha_nacimiento":"test"}`)
    req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
	
	expected := http.StatusCreated
	actual := rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }

    var m map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)

    if m["nombre"] != "test" {
        t.Errorf("Expected student name to be 'test'. Got '%v'", m["nombre"])
    }

	if m["dni"] != "12345678" {
		t.Errorf("Expected student dni to be '12345678'. Got '%v'", m["dni"])
	}
	if m["direccion"] != "test" {
		t.Errorf("Expected student direccion to be 'test'. Got '%v'", m["direccion"])
	}
	if m["fecha_nacimiento"] != "test" {
		t.Errorf("Expected student fecha_nacimiento to be 'test'. Got '%v'", m["fecha_nacimiento"])
	}
}


func TestGetStudent(t *testing.T) {

    req, _ := http.NewRequest("GET", "/product", nil)
	rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)
	
	expected := http.StatusOK
	actual := rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }


}




func TestUpdateProduct(t *testing.T) {

    req, _ := http.NewRequest("GET", "/product/12345678", nil)
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    var originalStudent map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &originalStudent)

    var jsonStr = []byte(`{"nombre":"test-updated", "dni":"12345678", "direccion":"test updated", "fecha_nacimiento":"test updated"}`)
    req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr = httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

	expected := http.StatusOK
	actual := rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
    var m map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)

    if m["dni"] != originalStudent["dni"] {
        t.Errorf("Expected the dni to remain the same (%v). Got %v", originalStudent["dni"], m["dni"])
    }
    if m["nombre"] == originalStudent["nombre"] {
        t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalStudent["name"], m["name"], m["name"])
    }
    if m["direccion"] == originalStudent["direccion"] {
		t.Errorf("Expected the direccion to change from '%v' to '%v'. Got '%v'", originalStudent["direccion"], m["direccion"], m["direccion"])
	}
	if m["fecha_nacimiento"] == originalStudent["fecha_nacimiento"] {
		t.Errorf("Expected the fecha_nacimiento to change from '%v' to '%v'. Got '%v'", originalStudent["fecha_nacimiento"], m["fecha_nacimiento"], m["fecha_nacimiento"])
	}

}


func TestDeleteProduct(t *testing.T) {
    req, _ := http.NewRequest("GET", "/product/12345678", nil)
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

	expected := http.StatusOK
	actual := rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }

    req, _ = http.NewRequest("DELETE", "/product/12345678", nil)
    rr = httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

	expected = http.StatusOK
	actual = rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }

    req, _ = http.NewRequest("GET", "/product/12345678", nil)
	rr = httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

	expected = http.StatusNotFound
	actual = rr.Code

	if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
