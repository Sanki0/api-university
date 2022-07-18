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

func createCourse(w http.ResponseWriter, r *http.Request) (int64, error) {

	var s models.Course
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	stmt, err := utils.DB.Prepare("INSERT INTO courses(nombre, descripcion, temas) VALUES(?,?,?)")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	result, err := stmt.Exec(s.Nombre, s.Descripcion, s.Temas)
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

func getCourses(w http.ResponseWriter) ([]*models.Course,error) {

	rows, err := utils.DB.Query("SELECT * FROM courses")
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var courses []*models.Course

	for rows.Next() {
		var course models.Course
		err = rows.Scan(&course.Nombre, &course.Descripcion, &course.Temas)
		utils.ChkError(err)
		courses = append(courses, &course)
	}

	if(courses == nil){
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusNotFound)
	}

	return courses,nil
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) (*models.Course, error){
	nombre := mux.Vars(r)["nombre"]

	query, err := utils.DB.Query("SELECT * FROM courses WHERE nombre = ?", nombre)

	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var s models.Course

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Descripcion, &s.Temas)
		utils.ChkError(err)
	}

	if s.Nombre == ""{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	}
	return &s,nil
}

func updateCourse(w http.ResponseWriter, r *http.Request) (int64,error) {
	var s models.Course
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	//prepare
	stmt, err := utils.DB.Prepare("UPDATE courses SET nombre = ?, descripcion = ?, temas = ? WHERE nombre = ?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(s.Nombre, s.Descripcion, s.Temas, s.Nombre)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	if(ro == 0){
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusNotFound)
	}

	if ro == 1 {
		w.Header().Set("Content-Type", "application/json")
   		w.WriteHeader(http.StatusNoContent)
	}


	return ro,nil

}

func deleteCourse(w http.ResponseWriter, r *http.Request) (int64,error) {
	var a models.Course
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	//prepare

	stmt, err := utils.DB.Prepare("DELETE FROM courses WHERE nombre = ?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(a.Nombre)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}
	ro, err := result.RowsAffected()
	utils.ChkError(err)

	if(ro == 0){
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusNotFound)
	}

	if ro == 1 {
		w.Header().Set("Content-Type", "application/json")
   		w.WriteHeader(http.StatusNoContent)
	}


	return ro,nil
}

/////

//CREATE
func CreateCoursePage(w http.ResponseWriter, r *http.Request) {
	id, err := createCourse(w, r)
	fmt.Fprintf(w, "Create Courses Page!\n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if err == nil {
		fmt.Fprintf(w, "Course created with id: %d", id)
	}

}

//READ
func ReadCoursesPage(w http.ResponseWriter, r *http.Request) {
	students,err := getCourses(w)
	fmt.Fprintf(w, "Courses Page: \n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if students == nil {
		fmt.Fprintf(w, "No courses found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func ReadCoursePage(w http.ResponseWriter, r *http.Request) {
	student,err := getSingleCourse(w, r)
	fmt.Fprintf(w, "Single Course Page: \n")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No course found")
	}
}

//UPDATE
func UpdateCoursePage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := updateCourse(w, r)
	fmt.Fprintf(w, "Update Course Page!\n")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Course updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Course not updated")
	}
}

//DELETE
func DeleteCoursePage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := deleteCourse(w, r)
	fmt.Fprintf(w, "Delete Course Page!\n")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Course deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Database not updated")
	}
}
