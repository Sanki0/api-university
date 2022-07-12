package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateAlumno(w http.ResponseWriter, r *http.Request) {

	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	chkError(err)

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	stmt, err := db.Prepare("INSERT INTO students(nombre, dni, direccion, fecha_nacimiento) VALUES(?,?,?,?)")
	chkError(err)

	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento)
	chkError(err)

	id, err := result.LastInsertId()
	chkError(err)
	fmt.Fprintf(w, "Student created with id: %d\n", id)
}

func GetStudents() []*Student {

	db := connectionDB()
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

func GetSingleStudent(w http.ResponseWriter, r *http.Request) *Student {
	var a Student
	err := json.NewDecoder(r.Body).Decode(&a)
	chkError(err)

	db := connectionDB()
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

func UpdateStudentPage(w http.ResponseWriter, r *http.Request) int64 {
	var s Student
	err := json.NewDecoder(r.Body).Decode(&s)
	chkError(err)

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare
	stmt, err := db.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE dni = ?")
	chkError(err)

	//execute
	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)
	chkError(err)

	ro, err := result.RowsAffected()
	chkError(err)

	return ro

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) int64 {
	var a Student
	err := json.NewDecoder(r.Body).Decode(&a)
	chkError(err)

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	//prepare

	stmt, err := db.Prepare("DELETE FROM students WHERE dni = ?")
	chkError(err)

	//execute
	result, err := stmt.Exec(a.Dni)
	chkError(err)

	ro, err := result.RowsAffected()
	chkError(err)

	return ro
}
