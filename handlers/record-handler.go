package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sanki0/api-university/models"
	"github.com/Sanki0/api-university/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func createRecord(w http.ResponseWriter, r *http.Request) error {

	var s models.Record
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	stmt, err := utils.DB.Prepare("INSERT INTO records (student, course, startdate, finishdate) VALUES (?,?,?,?)")
	utils.ChkError(err)

	result, err := stmt.Exec(s.Student, s.Course, s.Startdate, s.Finishdate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	utils.ChkError(err)
	fmt.Fprintf(w, "Record created with id: %d\n", id)

	return nil
}

func getRecords() []*models.Record {


	rows, err := utils.DB.Query("SELECT * FROM records")
	utils.ChkError(err)

	var records []*models.Record

	for rows.Next() {
		var record models.Record
		err = rows.Scan(&record.Student, &record.Course, &record.Startdate, &record.Finishdate)
		utils.ChkError(err)
		records = append(records, &record)
	}

	return records
}

func getSingleRecord(w http.ResponseWriter, r *http.Request) *models.Record {
	dni := mux.Vars(r)["dni"]
	course := mux.Vars(r)["course"]

	query, err := utils.DB.Query("SELECT * FROM records WHERE student = ? AND course = ?", dni, course)

	utils.ChkError(err)

	var record models.Record

	for query.Next() {
		err = query.Scan(&record.Student, &record.Course, &record.Startdate, &record.Finishdate)
		utils.ChkError(err)
	}
	return &record
}

func updateRecord(w http.ResponseWriter, r *http.Request) int64 {
	var s models.Record
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	//prepare
	stmt, err := utils.DB.Prepare("UPDATE records SET startdate = ?, finishdate = ? WHERE student = ? AND course=?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(s.Startdate, s.Finishdate, s.Student, s.Course)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro

}

func deleteRecord(w http.ResponseWriter, r *http.Request) int64 {
	var a models.Record
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	//prepare

	stmt, err := utils.DB.Prepare("DELETE FROM records WHERE student = ? AND course =?")
	utils.ChkError(err)

	//execute
	result, err := stmt.Exec(a.Student, a.Course)
	utils.ChkError(err)

	ro, err := result.RowsAffected()
	utils.ChkError(err)

	return ro
}

/////

//CREATE
func CreateRecordPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Record Page!\n")
	err := createRecord(w, r)
	if err != nil {
		fmt.Fprintf(w, "Error creating record")
	}
	if err == nil {
		fmt.Fprintf(w, "Record created")
	}

}

//READ
func ReadRecordsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Records Page: \n")
	records := getRecords()
	if records == nil {
		fmt.Fprintf(w, "No records found")
	}
	if records != nil {
		json.NewEncoder(w).Encode(records)
	}
}

func ReadRecordPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Single Record Page: \n")
	Record := getSingleRecord(w, r)
	if Record.Student != "" {
		json.NewEncoder(w).Encode(*Record)
	}
	if Record.Student == "" {
		fmt.Fprintf(w, "No Record found")
	}
}

//UPDATE
func UpdateRecordPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Record Page!\n")

	rowsAffected := updateRecord(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Record updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Record not updated")
	}
}

//DELETE
func DeleteRecordPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Record Page!\n")

	rowsAffected := deleteRecord(w, r)
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Record deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Record not deleted")
	}
}
