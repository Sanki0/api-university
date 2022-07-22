package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
    "github.com/Sanki0/api-university/src/models"
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

	a.Router = mux.NewRouter()  

	a.initializeRoutes()

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", nil))
}





//STUDENTS ROUTES

func (a *App) createStudent(w http.ResponseWriter, r *http.Request) {
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

func (a *App) getStudent(w http.ResponseWriter, r *http.Request) {
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

func (a *App) getStudents(w http.ResponseWriter, r *http.Request) {

    students, err :=  models.GetStudentsU(a.DB)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, students)
}



func (a *App) updateStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	dni := vars["dni"]
    var s models.Student
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&s); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    s.Dni = dni

    if err := s.UpdateStudent(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, s)
}

func (a *App) deleteStudent(w http.ResponseWriter, r *http.Request) {
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
    a.Router.HandleFunc("/student", a.createStudent).Methods("POST")
    a.Router.HandleFunc("/student", a.getStudents).Methods("GET")
    a.Router.HandleFunc("/student/{dni}", a.getStudent).Methods("GET")
    a.Router.HandleFunc("/student/{dni}", a.updateStudent).Methods("PUT")
    a.Router.HandleFunc("/student/{dni}", a.deleteStudent).Methods("DELETE")
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
}