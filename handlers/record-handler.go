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

func createRecord(w http.ResponseWriter, r *http.Request) (int64, error) {

	var s models.Record
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	stmt, err := utils.DB.Prepare("INSERT INTO records (student, course, startdate, finishdate) VALUES (?,?,?,?)")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	result, err := stmt.Exec(s.Student, s.Course, s.Startdate, s.Finishdate)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())

		utils.RespondWithError(w, http.StatusConflict, err.Error())
		return -1,err
	}

	id, err := result.LastInsertId()
	utils.ChkError(err)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

	return id, nil
}

func getRecords(w http.ResponseWriter) ([]*models.Record, error) {


	rows, err := utils.DB.Query("SELECT * FROM records")
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var records []*models.Record

	for rows.Next() {
		var record models.Record
		err = rows.Scan(&record.Student, &record.Course, &record.Startdate, &record.Finishdate)
		utils.ChkError(err)
		records = append(records, &record)
	}

	if(records == nil){
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusNotFound)
	}

	return records,nil
}

func getSingleRecord(w http.ResponseWriter, r *http.Request) (*models.Record, error) {
	dni := mux.Vars(r)["dni"]
	course := mux.Vars(r)["course"]

	query, err := utils.DB.Query("SELECT * FROM records WHERE student = ? AND course = ?", dni, course)
	if err != nil {
		err = fmt.Errorf("Error executing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return nil,err
	}

	var record models.Record

	for query.Next() {
		err = query.Scan(&record.Student, &record.Course, &record.Startdate, &record.Finishdate)
		utils.ChkError(err)
	}

	if(record.Course == ""){
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusNotFound)
	}

	return &record,nil
}

func updateRecord(w http.ResponseWriter, r *http.Request) (int64, error) {
	var s models.Record
	err := json.NewDecoder(r.Body).Decode(&s)
	utils.ChkError(err)

	//prepare
	stmt, err := utils.DB.Prepare("UPDATE records SET startdate = ?, finishdate = ? WHERE student = ? AND course=?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(s.Startdate, s.Finishdate, s.Student, s.Course)
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

func deleteRecord(w http.ResponseWriter, r *http.Request) (int64,error) {
	var a models.Record
	err := json.NewDecoder(r.Body).Decode(&a)
	utils.ChkError(err)

	//prepare

	stmt, err := utils.DB.Prepare("DELETE FROM records WHERE student = ? AND course =?")
	if err != nil {
		err = fmt.Errorf("Error preparing query\n %q", err.Error())
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return -1,err
	}

	//execute
	result, err := stmt.Exec(a.Student, a.Course)
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
func CreateRecordPage(w http.ResponseWriter, r *http.Request) {
	id, err := createRecord(w, r)
	
	fmt.Fprintf(w, "Create Record Page!\n")
	
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	if err == nil {
		fmt.Fprintf(w, "Record created with id: %d", id)
	}

}

//READ
func ReadRecordsPage(w http.ResponseWriter, r *http.Request) {
	records,err := getRecords(w)
	
	fmt.Fprintf(w, "Records Page: \n")
	
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	if records == nil {
		fmt.Fprintf(w, "No records found")
	}
	if records != nil {
		json.NewEncoder(w).Encode(records)
	}
}

func ReadRecordPage(w http.ResponseWriter, r *http.Request) {
	Record,err := getSingleRecord(w, r)
	
	fmt.Fprintf(w, "Single Record Page: \n")
	
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	if Record.Student != "" {
		json.NewEncoder(w).Encode(*Record)
	}
	if Record.Student == "" {
		fmt.Fprintf(w, "No Record found")
	}
}

//UPDATE
func UpdateRecordPage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := updateRecord(w, r)
	
	fmt.Fprintf(w, "Update Record Page!\n")

	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Record updated")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Record not updated")
	}
}

//DELETE
func DeleteRecordPage(w http.ResponseWriter, r *http.Request) {
	rowsAffected,err := deleteRecord(w, r)
	
	fmt.Fprintf(w, "Delete Record Page!\n")

	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	if rowsAffected > 0 {
		fmt.Fprintf(w, "Record deleted")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Record not deleted")
	}
}
