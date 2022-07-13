package models

type Course struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Temas       string `json:"temas"`
}
