package models

import (
    "database/sql"
)


type Course struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Temas       string `json:"temas"`
}


func (c *Course) GetCourse (db *sql.DB) error {
  return db.QueryRow("SELECT * FROM courses WHERE nombre = ?", 
  					c.Nombre).Scan(&c.Nombre, &c.Descripcion, &c.Temas)
}

func (c *Course) UpdateCourse(db *sql.DB) error {
    _, err :=
	db.Exec("UPDATE courses SET nombre = ?, descripcion = ?, temas = ? WHERE nombre = ?", 
				c.Nombre, c.Descripcion, c.Temas, c.Nombre)

	return err
}

func (c *Course)  DeleteCourse(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM courses WHERE nombre = ?", c.Nombre)

    return err
}

func (c *Course)  CreateCourse(db *sql.DB) error {
	_,err := db.Exec("INSERT INTO courses(nombre, descripcion, temas) VALUES(?,?,?)", 
						c.Nombre, c.Descripcion, c.Temas)

    if err != nil {
        return err
    }

    return nil
}

func GetCoursesU(db *sql.DB) ([]Course, error) {
	rows, err := db.Query("SELECT * FROM courses")


    if err != nil {
        return nil, err
    }

    defer rows.Close()

    courses := []Course{}

    for rows.Next() {
        var c Course
        if err := rows.Scan(&c.Nombre, &c.Descripcion, &c.Temas); err != nil {
			return nil, err
		}
        courses = append(courses, c)
    }

    return courses, nil
}