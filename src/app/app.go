package app

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {

	// Initialize the database connection
	var err error
	a.DB, err = sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		log.Fatal(err)
	}

    a.DB.SetMaxIdleConns(25)
	a.DB.SetMaxOpenConns(25)
	a.DB.SetConnMaxLifetime(5*time.Minute)
	a.DB.SetConnMaxIdleTime(3*time.Minute)

	a.Router = mux.NewRouter()  

	a.initializeRoutes()

}

func (a *App) Run(addr string) {
    http.Handle("/", a.Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/student", a.CreateStudent).Methods("POST")
    a.Router.HandleFunc("/student", a.GetStudents).Methods("GET")
    a.Router.HandleFunc("/student", a.UpdateStudent).Methods("PUT")
    a.Router.HandleFunc("/student/{dni}", a.GetStudent).Methods("GET")
    a.Router.HandleFunc("/student/{dni}", a.DeleteStudent).Methods("DELETE")

    a.Router.HandleFunc("/course", a.CreateCourse).Methods("POST")
    a.Router.HandleFunc("/course", a.GetCourses).Methods("GET")
    a.Router.HandleFunc("/course", a.UpdateCourse).Methods("PUT")
    a.Router.HandleFunc("/course/{nombre}", a.GetCourse).Methods("GET")
    a.Router.HandleFunc("/course/{nombre}", a.DeleteCourse).Methods("DELETE")

	a.Router.HandleFunc("/record", a.CreateRecord).Methods("POST")
	a.Router.HandleFunc("/record", a.GetRecords).Methods("GET")
	a.Router.HandleFunc("/record", a.UpdateRecord).Methods("PUT")
	a.Router.HandleFunc("/record/{student}/{course}", a.GetRecord).Methods("GET")
	a.Router.HandleFunc("/record/{student}/{course}", a.DeleteRecord).Methods("DELETE")

}


