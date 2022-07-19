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

func createAlumno(w http.ResponseWriter, r *http.Request) (int64, error) {

	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)


	stmt, err := utils.DB.Prepare("INSERT INTO students (nombre, dni, direccion,fecha_nacimiento) VALUES (?,?,?,?)")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusConflict, err.Error())
		return -1,err
	}

	id, err := result.LastInsertId()
	utils.ChkError(err)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

	return id,nil
}

func getStudents(w http.ResponseWriter) ([]*models.Student, error) {

	rows, err := utils.DB.Query("SELECT * FROM students")
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var students []*models.Student

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.Nombre, &student.Dni, &student.Direccion, &student.Fecha_nacimiento)
		utils.ChkError(err)
		students = append(students, &student)
	}

	w.Header().Set("Content-Type", "application/json")
	if students == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	return students,nil
}

func getSingleStudent(w http.ResponseWriter, r *http.Request) (*models.Student, error) {
	
	dni := mux.Vars(r)["dni"]
	query, err := utils.DB.Query("SELECT * FROM students WHERE dni = ?", dni)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var s models.Student

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
		utils.ChkError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if(s.Dni == ""){
    	w.WriteHeader(http.StatusNotFound)
	}

	return &s,nil
}

func updateStudent(w http.ResponseWriter, r *http.Request) (int64,error) {
	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	//prepare
	stmt, err := utils.DB.Prepare("UPDATE students SET nombre = ?, dni = ?, direccion = ?, fecha_nacimiento = ? WHERE dni = ?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	w.Header().Set("Content-Type", "application/json")
	if(ro == 0){
    	w.WriteHeader(http.StatusNotFound)
	}

	if ro == 1 {
   		w.WriteHeader(http.StatusNoContent)
	}

	return ro,nil

}

func deleteStudent(w http.ResponseWriter, r *http.Request) (int64,error) {
	var a models.Student
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)
	
	//prepare
	stmt, err := utils.DB.Prepare("DELETE FROM students WHERE dni = ?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(a.Dni)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	ro, err := result.RowsAffected()
	utils.ChkError(err)
	
	w.Header().Set("Content-Type", "application/json")

	if(ro == 0){
    	w.WriteHeader(http.StatusNotFound)
	}

	if ro == 1 {
   		w.WriteHeader(http.StatusNoContent)
	}


	return ro,nil
}

/////

//CREATE
func CreateStudentPage(w http.ResponseWriter, r *http.Request) {
	id,err := createAlumno(w, r)
	fmt.Fprintf(w, "Create Student Page!\n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err == nil {
		fmt.Fprintf(w, "Student created with id: %d\n", id)
	}

}

//READ
func ReadStudentsPage(w http.ResponseWriter, r *http.Request) {
	students,err := getStudents(w)
	fmt.Fprintf(w, "Students Page: \n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if students == nil {
		fmt.Fprintf(w, "No students found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func ReadStudentPage(w http.ResponseWriter, r *http.Request) {
	student,err := getSingleStudent(w, r)
	fmt.Fprintf(w, "Single Student Page: \n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No student found")
	}
}

//UPDATE
func UpdateStudentPage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := updateStudent(w, r)
	fmt.Fprintf(w, "Update Student Page!\n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not updated")
	}
}

//DELETE
func DeleteStudentPage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := deleteStudent(w, r)
	fmt.Fprintf(w, "Delete Student Page!\n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Student deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Student not deleted")
	}
}
