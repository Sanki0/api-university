package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Sanki0/api-university/src/app/utils"
	"github.com/Sanki0/api-university/src/models"
	"github.com/gorilla/mux"
)

//STUDENTS ROUTES

func (a *App) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
	if err := s.CreateStudent(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, s)
}

func (a *App) GetStudent(w http.ResponseWriter, r *http.Request) {
    dni := mux.Vars(r)["dni"]

    s := models.Student{Dni: dni}
    if err := s.GetStudent(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            utils.RespondWithError(w, http.StatusNotFound, "Student not found")
        default:
            utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, s)
}

func (a *App) GetStudents(w http.ResponseWriter, r *http.Request) {

    students, err :=  models.GetStudentsU(a.DB)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, students)
}



func (a *App) UpdateStudent(w http.ResponseWriter, r *http.Request) {
    var s models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
    err := s.UpdateStudent(a.DB);
    if  err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, s)
}

func (a *App) DeleteStudent(w http.ResponseWriter, r *http.Request) {
    dni := mux.Vars(r)["dni"]

    p := models.Student{Dni: dni}

    if err := p.DeleteStudent(a.DB); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}