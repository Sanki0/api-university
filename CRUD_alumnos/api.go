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
	Nombre string `json:"nombre"`
	Dni string `json:"dni"`
	Direccion string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}

func chkError(err error) {
	if err != nil {
		panic(err)
	}
}

func PingDb(db *sql.DB){
	err := db.Ping()
	chkError(err)
}

func connectionDB() *sql.DB{
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	return db
}



func createAlumno(w http.ResponseWriter, r *http.Request) {

	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	chkError(err)

	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	stmt, err := db.Prepare("INSERT INTO students(nombre, dni, direccion, fecha_nacimiento) VALUES(?,?,?,?)")
	chkError(err)
	
	result,err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento)
	chkError(err)

	id, err := result.LastInsertId()
	chkError(err)
	fmt.Fprintf(w, "Student created with id: %d\n", id)
}

func getStudents() []*Student {
	
	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	rows, err := db.Query("SELECT * FROM students")
	chkError(err)

	var students []*Student

	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Nombre, &student.Dni, &student.Direccion, &student.Fecha_nacimiento)
		chkError(err)
		students = append(students, &student)
	}

	return students
}


func getSingleStudent(w http.ResponseWriter, r *http.Request) *Student{
    var a Student
	err := json.NewDecoder(r.Body).Decode(&a)
	chkError(err)

    db :=connectionDB()
	defer db.Close()
	PingDb(db)

	query, err := db.Query("SELECT * FROM students WHERE dni = ?", a.Dni)

	chkError(err)

	var s Student

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		chkError(err)
	}
	return &s
}

func updateStudentPage(w http.ResponseWriter, r *http.Request) int64 {
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	chkError(err)

	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare
	stmt, err:= db.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE dni = ?")
	chkError(err)

	//execute
	result,err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)
	chkError(err)

	ro,err := result.RowsAffected()
	chkError(err)

	return ro

}


func deleteStudent(w http.ResponseWriter, r *http.Request) int64{
	var a Student
	err := json.NewDecoder(r.Body).Decode(&a)
	chkError(err)
	
	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare

	stmt, err:= db.Prepare("DELETE FROM students WHERE dni = ?")
	chkError(err)

	//execute
	result,err := stmt.Exec(a.Dni)
	chkError(err)

	ro,err := result.RowsAffected()
	chkError(err)

	return ro
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

//CREATE
func createPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Page!\n")
	createAlumno(w,r)
	fmt.Fprintf(w, "Student created")

}

//READ
func studentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Students Page: \n")
	students := getStudents()
	if students == nil{
		fmt.Fprintf(w, "No students found")
	}
	if students != nil{
		json.NewEncoder(w).Encode(students)
	}
}

func singleStudentPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Single Student Page: \n")
	student := getSingleStudent(w, r)
	if student.Nombre != ""{
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == ""{
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func updatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Page!\n")

	rowsAffected := updateStudentPage(w, r)
	if rowsAffected > 0{
		fmt.Fprintf(w, "Student updated")
	}
	if rowsAffected == 0{
		fmt.Fprintf(w, "Student not updated")
	}
}


//DELETE
func deleteStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Page!\n")
	
	rowsAffected := deleteStudent(w,r);
	if rowsAffected > 0{
		fmt.Fprintf(w, "Student deleted")
	}
	if rowsAffected == 0{
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