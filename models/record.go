package models

type Record struct {
	Student    string `json:"student"`
	Course     string `json:"course"`
	Startdate  string `json:"startdate"`
	Finishdate string `json:"finishdate"`
}
