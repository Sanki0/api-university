package main

import (
	"github.com/Sanki0/api-university/handlers"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

//CREATE
func createPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Page!\n")
	handlers.CreateAlumno(w, r)
	fmt.Fprintf(w, "Student created")

}

//READ
func studentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Students Page: \n")
	students := handlers.GetStudents()
	if students == nil {
		fmt.Fprintf(w, "No students found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func singleStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Student Page: \n")
	student := handlers.GetSingleStudent(w, r)
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func updatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Page!\n")

	rowsAffected := handlers.UpdateStudentPage(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not updated")
	}
}

//DELETE
func deleteStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Page!\n")

	rowsAffected := handlers.DeleteStudent(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not deleted")
	}
}

func main() {

	r := mux.NewRouter()

	//CREATE STUDENT
	r.HandleFunc("/student", createPage).Methods("POST")

	//READ ALL STUDENTS
	r.HandleFunc("/students", studentPage).Methods("GET")

	//READ SINGLE STUDENT
	r.HandleFunc("/student", singleStudentPage).Methods("GET")

	//UPDATE STUDENT
	r.HandleFunc("/student", updatePage).Methods("PUT")
	//DELETE STUDENT
	r.HandleFunc("/student", deleteStudentPage).Methods("DELETE")

	//HOME PAGE

	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
