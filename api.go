package main

import (
	"github.com/Sanki0/api-university/handlers"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {

	r := mux.NewRouter()

	//STUDENTS
	//CREATE STUDENT
	r.HandleFunc("/student", handlers.CreateStudentPage).Methods("POST")

	//READ ALL STUDENTS
	r.HandleFunc("/students", handlers.ReadStudentsPage).Methods("GET")

	//READ SINGLE STUDENT
	r.HandleFunc("/student", handlers.ReadStudentPage).Methods("GET")

	//UPDATE STUDENT
	r.HandleFunc("/student", handlers.UpdateStudentPage).Methods("PUT")

	//DELETE STUDENT
	r.HandleFunc("/student", handlers.DeleteStudentPage).Methods("DELETE")

	//COURSES
	//CREATE COURSE
	r.HandleFunc("/course", handlers.CreateCoursePage).Methods("POST")

	//READ ALL COURSES
	r.HandleFunc("/courses", handlers.ReadCoursesPage).Methods("GET")

	//READ SINGLE COURSE
	r.HandleFunc("/course", handlers.ReadCoursePage).Methods("GET")

	//UPDATE COURSE
	r.HandleFunc("/course", handlers.UpdateCoursePage).Methods("PUT")

	//DELETE COURSE
	r.HandleFunc("/course", handlers.DeleteCoursePage).Methods("DELETE")

	//HOME PAGE
	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
