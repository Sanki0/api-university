package handlers

import (
	"github.com/Sanki0/api-university/models"
	"github.com/Sanki0/api-university/utils"

	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateAlumno(w http.ResponseWriter, r *http.Request) {

	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	stmt, err := db.Prepare("INSERT INTO students(nombre, dni, direccion, fecha_nacimiento) VALUES(?,?,?,?)")
	utils.ChkError(err)

	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento)
	utils.ChkError(err)

	id, err := result.LastInsertId()
	utils.ChkError(err)
	fmt.Fprintf(w, "Student created with id: %d\n", id)
}

func GetStudents() []*models.Student {

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	rows, err := db.Query("SELECT * FROM students")
	utils.ChkError(err)

	var students []*models.Student

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.Nombre, &student.Dni, &student.Direccion, &student.Fecha_nacimiento)
		utils.ChkError(err)
		students = append(students, &student)
	}

	return students
}

func GetSingleStudent(w http.ResponseWriter, r *http.Request) *models.Student {
	var a models.Student
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	query, err := db.Query("SELECT * FROM students WHERE dni = ?", a.Dni)

	utils.ChkError(err)

	var s models.Student

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		utils.ChkError(err)
	}
	return &s
}

func UpdateStudentPage(w http.ResponseWriter, r *http.Request) int64 {
	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	//prepare
	stmt, err := db.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE dni = ?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) int64 {
	var a models.Student
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	//prepare

	stmt, err := db.Prepare("DELETE FROM students WHERE dni = ?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(a.Dni)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro
}

/////

//CREATE
func CreatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Page!\n")
	CreateAlumno(w, r)
	fmt.Fprintf(w, "Student created")

}

//READ
func StudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Students Page: \n")
	students := GetStudents()
	if students == nil {
		fmt.Fprintf(w, "No students found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func SingleStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Student Page: \n")
	student := GetSingleStudent(w, r)
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func UpdatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Page!\n")

	rowsAffected := UpdateStudentPage(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not updated")
	}
}

//DELETE
func DeleteStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Page!\n")

	rowsAffected := DeleteStudent(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not deleted")
	}
}
