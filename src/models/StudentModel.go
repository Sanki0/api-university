package models

import (
    "database/sql"
)


type Student struct {
	Nombre           string `json:"nombre"`
	Dni              string `json:"dni"`
	Direccion        string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}


func (s *Student) GetStudent(db *sql.DB) error {
  return db.QueryRow("SELECT * FROM students WHERE dni = ?", 
  					s.Dni).Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento)
}

func (s *Student)  UpdateStudent(db *sql.DB) error {
	_, err :=
	db.Exec("UPDATE products SET Nombre=?, Dni=?, Direccion=?, Fecha_nacimiento=? WHERE dni=?", 
				s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento, s.Dni)

	return err
}

func (s *Student)  DeleteStudent(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM students WHERE dni = ?", s.Dni)

    return err
}

func (s *Student)  CreateStudent(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO students(Nombre, Dni, Direccion, Fecha_nacimiento) VALUES(?, ?, ?, ?) RETURNING dni", 
						s.Nombre, s.Dni, s.Direccion, s.Fecha_nacimiento).Scan(&s.Dni)

    if err != nil {
        return err
    }

    return nil
}

func GetStudentsU(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM students")


    if err != nil {
        return nil, err
    }

    defer rows.Close()

    students := []Student{}

    for rows.Next() {
        var s Student
        if err := rows.Scan(&s.Nombre, &s.Dni, &s.Direccion, &s.Fecha_nacimiento); err != nil {
			return nil, err
		}
        students = append(students, s)
    }

    return students, nil
}