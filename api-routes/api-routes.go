package apiroutes

import (
	"fmt"
	"log"
	"github.com/Sanki0/api-university/handlers"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}


func InitApiRoutes(r *mux.Router) {
//STUDENTS
	//CREATE STUDENT
	r.HandleFunc("/student", handlers.CreateStudentPage).Methods("POST")

	//READ ALL STUDENTS
	r.HandleFunc("/student", handlers.ReadStudentsPage).Methods("GET")

	//READ SINGLE STUDENT
	r.HandleFunc("/student/{dni}", handlers.ReadStudentPage).Methods("GET")

	//UPDATE STUDENT
	r.HandleFunc("/student", handlers.UpdateStudentPage).Methods("PUT")

	//DELETE STUDENT
	r.HandleFunc("/student", handlers.DeleteStudentPage).Methods("DELETE")

//COURSES
	//CREATE COURSE
	r.HandleFunc("/course", handlers.CreateCoursePage).Methods("POST")

	//READ ALL COURSES
	r.HandleFunc("/course", handlers.ReadCoursesPage).Methods("GET")

	//READ SINGLE COURSE
	r.HandleFunc("/course/{nombre}", handlers.ReadCoursePage).Methods("GET")

	//UPDATE COURSE
	r.HandleFunc("/course", handlers.UpdateCoursePage).Methods("PUT")

	//DELETE COURSE
	r.HandleFunc("/course", handlers.DeleteCoursePage).Methods("DELETE")

//RECORDS
	//CREATE RECORD
	r.HandleFunc("/record", handlers.CreateRecordPage).Methods("POST")

	//READ ALL RECORDS
	r.HandleFunc("/record", handlers.ReadRecordsPage).Methods("GET")

	//READ SINGLE RECORD
	r.HandleFunc("/record/{dni}/{course}", handlers.ReadRecordPage).Methods("GET")

	//UPDATE RECORD
	r.HandleFunc("/record", handlers.UpdateRecordPage).Methods("PUT")

	//DELETE RECORD
	r.HandleFunc("/record", handlers.DeleteRecordPage).Methods("DELETE")

	//HOME PAGE
	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}