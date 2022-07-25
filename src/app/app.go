package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Sanki0/api-university/src/models"
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





//STUDENTS ROUTES

func (a *App) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
	if err := s.CreateStudent(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, s)
}

func (a *App) GetStudent(w http.ResponseWriter, r *http.Request) {
    dni := mux.Vars(r)["dni"]

    s := models.Student{Dni: dni}
    if err := s.GetStudent(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Student not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    respondWithJSON(w, http.StatusOK, s)
}

func (a *App) GetStudents(w http.ResponseWriter, r *http.Request) {

    students, err :=  models.GetStudentsU(a.DB)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, students)
}



func (a *App) UpdateStudent(w http.ResponseWriter, r *http.Request) {
    var s models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
    err := s.UpdateStudent(a.DB);
    if  err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, s)
}

func (a *App) DeleteStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	dni := vars["dni"]

    p := models.Student{Dni: dni}
    if err := p.DeleteStudent(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/student", a.CreateStudent).Methods("POST")
    a.Router.HandleFunc("/student", a.GetStudents).Methods("GET")
    a.Router.HandleFunc("/student/{dni}", a.GetStudent).Methods("GET")
    a.Router.HandleFunc("/student", a.UpdateStudent).Methods("PUT")
    a.Router.HandleFunc("/student/{dni}", a.DeleteStudent).Methods("DELETE")
}


//RESPONSES
func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
    //log.Println(string(response))
}