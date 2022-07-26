package models

import "database/sql"

type Record struct {
	Student    string `json:"student"`
	Course     string `json:"course"`
	Startdate  string `json:"startdate"`
	Finishdate string `json:"finishdate"`
}



func (r * Record) GetRecord (db *sql.DB) error {
	return db.QueryRow("SELECT * FROM records WHERE student = ? AND course = ?", 
					r.Student, r.Course).Scan(&r.Student, &r.Course, &r.Startdate, &r.Finishdate)
						
  }
  
  func (r * Record) UpdateRecord(db *sql.DB) error {
	  _, err :=
	  db.Exec("UPDATE records SET student = ?, course = ?, startdate = ?, finishdate = ? WHERE student = ? AND course = ?",
	 			r.Student, r.Course, r.Startdate, r.Finishdate, r.Student, r.Course)			
  
	  return err
  }
  
  func (r * Record)  DeleteRecord(db *sql.DB) error {
	  _, err := db.Exec("DELETE FROM records WHERE student = ? AND course = ?", r.Student, r.Course)
  
	  return err
  }
  
  func (r * Record)  CreateRecord(db *sql.DB) error {
	  _,err := db.Exec("INSERT INTO records(student, course, startdate, finishdate) VALUES(?, ?, ?, ?)",
	  				r.Student, r.Course, r.Startdate, r.Finishdate)
  
	  if err != nil {
		  return err
	  }
  
	  return nil
  }
  
  func GetRecordsU(db *sql.DB) ([]Record, error) {
	  rows, err := db.Query("SELECT * FROM records")
  
  
	  if err != nil {
		  return nil, err
	  }
  
	  defer rows.Close()
  
	  records := []Record{}
  
	  for rows.Next() {
		  var r Record
		  if err := rows.Scan(&r.Student, &r.Course, &r.Startdate, &r.Finishdate); err != nil {
			  return nil, err
		  }
		  records = append(records, r)
	  }
  
	  return records, nil
  }