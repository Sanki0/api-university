package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)

type Student struct {
	Id   string    `json:"id"`
	Nombre string `json:"name"`
	Dni string `json:"dni"`
	Direccion string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}

func getStudents() []*Student {
	// Open up our database connection.
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM students")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var students []*Student
	for results.Next() {
		var s Student
		// for each row, scan the result into our tag composite object
		err = results.Scan(&s.Id, &s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		students = append(students, &s)
	}

	return students
}


func getSingleStudent(w http.ResponseWriter, r *http.Request) *Student{
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query, err := db.Query("SELECT * FROM students WHERE id = ?", vars["id"])
	if err != nil {
		panic(err.Error())
	}
	var s Student
	for query.Next() {
		err = query.Scan(&s.Id, &s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		if err != nil {
			panic(err.Error())
		}
	}
	return &s
}

func deleteStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query, err := db.Query("DELETE FROM students WHERE id = ?", vars["id"])
	if err != nil {
		panic(err.Error())
	}

	var s Student
	for query.Next() {
		err = query.Scan(&s.Id, &s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		if err != nil {
			panic(err.Error())
		}
	}
	
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func studentPage(w http.ResponseWriter, r *http.Request) {
	students := getStudents()
	json.NewEncoder(w).Encode(students)
}

func singleStudentPage(w http.ResponseWriter, r *http.Request){
	student := getSingleStudent(w, r)
	json.NewEncoder(w).Encode(student)
}

func deleteStudentPage(w http.ResponseWriter, r *http.Request) {
	deleteStudent(w,r);
	fmt.Println("DELETED STUDENT")
}


func main() {
	
	r := mux.NewRouter()

	//CREATE STUDENT
	
	//READ ALL STUDENTS
	r.HandleFunc("/students", studentPage)
	
	//READ SINGLE STUDENT
	r.HandleFunc("/student/{id}", singleStudentPage)
	
	//UPDATE STUDENT

	//DELETE STUDENT
	r.HandleFunc("/delete/{id}", deleteStudent)



	
	//HOME PAGE

	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}