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
	vars := mux.Vars(r)

	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	stmt, err := db.Prepare("INSERT INTO students(nombre, dni, direccion, fecha_nacimiento) VALUES(?,?,?,?)")
	chkError(err)
	
	result,err := stmt.Exec(vars["nombre"], vars["dni"], vars["direccion"], vars["fecha_nacimiento"])
	chkError(err)

	id,err := result.LastInsertId()
	chkError(err)
	
	fmt.Println("Last inserted ID is:", id)
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
		err = rows.Scan(&student.Id, &student.Nombre, &student.Dni, &student.Direccion, &student.Fecha_nacimiento)
		chkError(err)
		students = append(students, &student)
	}

	return students
}


func getSingleStudent(w http.ResponseWriter, r *http.Request) *Student{
    vars := mux.Vars(r)

    db :=connectionDB()
	defer db.Close()
	PingDb(db)

	query, err := db.Query("SELECT * FROM students WHERE id = ?", vars["id"])
	chkError(err)

	var s Student
	for query.Next() {
		err = query.Scan(&s.Id, &s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		chkError(err)
	}

	return &s
}

func updateStudentPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare
	stmt, err:= db.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE id = ?")
	chkError(err)

	//execute
	result,err := stmt.Exec(vars["nombre"], vars["dni"], vars["direccion"], vars["fecha_nacimiento"], vars["id"])
	chkError(err)

	ro,err := result.RowsAffected()
	chkError(err)

	fmt.Println(ro)
}


func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	db :=connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare

	stmt, err:= db.Prepare("DELETE FROM students WHERE id = ?")
	chkError(err)

	//execute
	result,err := stmt.Exec(vars["id"])
	chkError(err)

	ro,err := result.RowsAffected()
	chkError(err)

	fmt.Println(ro)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

//CREATE
func createPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Page!")
	createAlumno(w,r)
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
	if student.Id != ""{
		json.NewEncoder(w).Encode(student)
	}
	if student.Id == ""{
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func updatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Page!")
	updateStudentPage(w, r)
}


//DELETE
func deleteStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Page!")
	deleteStudent(w,r);
}


func main() {
	
	r := mux.NewRouter()

	//CREATE STUDENT
	r.HandleFunc("/student/{nombre}/{dni}/{direccion}/{fecha_nacimiento}", createPage)
	
	//READ ALL STUDENTS
	r.HandleFunc("/students", studentPage).Methods("GET")
	
	//READ SINGLE STUDENT
	r.HandleFunc("/student/{id}", singleStudentPage)
	
	//UPDATE STUDENT
	r.HandleFunc("/student/{id}/{nombre}/{dni}/{direccion}/{fecha_nacimiento}", updatePage)

	//DELETE STUDENT
	r.HandleFunc("/delete/{id}", deleteStudent)
	
	//HOME PAGE

	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}