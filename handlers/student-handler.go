package handlers

import (
	"github.com/Sanki0/api-university/models"
	"github.com/Sanki0/api-university/utils"
	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createAlumno(w http.ResponseWriter, r *http.Request) error {

	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)


	stmt, err := utils.DB.Prepare("INSERT INTO students (nombre, dni, direccion,fecha_nacimiento) VALUES (?,?,?,?)")
	utils.ChkError(err)

	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	utils.ChkError(err)
	fmt.Fprintf(w, "Student created with id: %d\n", id)
	return nil
}

func getStudents() []*models.Student {

	rows, err := utils.DB.Query("SELECT * FROM students")
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

func getSingleStudent(w http.ResponseWriter, r *http.Request) *models.Student {
	
	dni := mux.Vars(r)["dni"]
	query, err := utils.DB.Query("SELECT * FROM students WHERE dni = ?", dni)
	utils.ChkError(err)

	var s models.Student

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		utils.ChkError(err)
	}
	return &s
}

func updateStudent(w http.ResponseWriter, r *http.Request) int64 {
	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	//prepare
	stmt, err := utils.DB.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE dni = ?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro

}

func deleteStudent(w http.ResponseWriter, r *http.Request) int64 {
	var a models.Student
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)
	
	//prepare
	stmt, err := utils.DB.Prepare("DELETE FROM students WHERE dni = ?")
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
func CreateStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Student Page!\n")
	err := createAlumno(w, r)
	if err != nil {
		fmt.Fprintf(w, "Student not created")
	}
	if err == nil {
		fmt.Fprintf(w, "Student created")
	}

}

//READ
func ReadStudentsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Students Page: \n")
	students := getStudents()
	if students == nil {
		fmt.Fprintf(w, "No students found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func ReadStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Student Page: \n")
	student := getSingleStudent(w, r)
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func UpdateStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Student Page!\n")

	rowsAffected := updateStudent(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not updated")
	}
}

//DELETE
func DeleteStudentPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Student Page!\n")

	rowsAffected := deleteStudent(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not deleted")
	}
}
