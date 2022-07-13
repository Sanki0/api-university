package handlers

import (
	"github.com/Sanki0/api-university/models"
	"github.com/Sanki0/api-university/utils"

	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createCourse(w http.ResponseWriter, r *http.Request) error {

	var s models.Course
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	stmt, err := db.Prepare("INSERT INTO courses(nombre, descripcion, temas) VALUES(?,?,?)")
	utils.ChkError(err)

	result, err := stmt.Exec(s.Nombre, s.Descripcion, s.Temas)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	utils.ChkError(err)
	fmt.Fprintf(w, "Course created with id: %d\n", id)
	return nil
}

func getCourses() []*models.Course {

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	rows, err := db.Query("SELECT * FROM courses")
	utils.ChkError(err)

	var courses []*models.Course

	for rows.Next() {
		var course models.Course
		err = rows.Scan(&course.Nombre, &course.Descripcion, &course.Temas)
		utils.ChkError(err)
		courses = append(courses, &course)
	}

	return courses
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) *models.Course {
	var a models.Student
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	query, err := db.Query("SELECT * FROM courses WHERE nombre = ?", a.Nombre)

	utils.ChkError(err)

	var s models.Course

	for query.Next() {
		err = query.Scan(&s.Nombre, &s.Descripcion, &s.Temas)
		utils.ChkError(err)
	}
	return &s
}

func updateCourse(w http.ResponseWriter, r *http.Request) int64 {
	var s models.Course
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	//prepare
	stmt, err := db.Prepare("UPDATE courses SET nombre = ?, descripcion = ?, temas = ? WHERE nombre = ?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(s.Nombre, s.Descripcion, s.Temas, s.Nombre)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro

}

func deleteCourse(w http.ResponseWriter, r *http.Request) int64 {
	var a models.Course
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	db := utils.ConnectionDB()
	defer db.Close()
	utils.PingDb(db)

	//prepare

	stmt, err := db.Prepare("DELETE FROM courses WHERE nombre = ?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(a.Nombre)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro
}

/////

//CREATE
func CreateCoursePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Courses Page!\n")
	err := createCourse(w, r)
	if err != nil {
		fmt.Fprintf(w, "Reapeted name")
	}
	if err == nil {
		fmt.Fprintf(w, "Course created")
	}

}

//READ
func ReadCoursesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Courses Page: \n")
	students := getCourses()
	if students == nil {
		fmt.Fprintf(w, "No courses found")
	}
	if students != nil {
		json.NewEncoder(w).Encode(students)
	}
}

func ReadCoursePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Course Page: \n")
	student := getSingleCourse(w, r)
	if student.Nombre != "" {
		json.NewEncoder(w).Encode(*student)
	}
	if student.Nombre == "" {
		fmt.Fprintf(w, "No course found")
	}
}

//UPDATE
func UpdateCoursePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Course Page!\n")

	rowsAffected := updateCourse(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Course updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Course not updated")
	}
}

//DELETE
func DeleteCoursePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Course Page!\n")

	rowsAffected := deleteCourse(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Course deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Database not updated")
	}
}
