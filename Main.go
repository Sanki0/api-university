package main

import (
	"fmt"
	//"log"
	"net/http"
	"text/template"
)

var plantilla= template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", initHandler)// acceder a la funcion inicio
	http.HandleFunc("/crear", createHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}


func initHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World!") //mostrar un saludo
 	plantilla.ExecuteTemplate(w, "inicio", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World!") //mostrar un saludo
 	plantilla.ExecuteTemplate(w, "crear", nil)
}