package models

type Student struct {
	Nombre           string `json:"nombre"`
	Dni              string `json:"dni"`
	Direccion        string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}
