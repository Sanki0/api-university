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

func (a *App) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var c models.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
	if err := c.CreateCourse(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, c)
}

func (a *App) GetCourse(w http.ResponseWriter, r *http.Request) {
    nombre := mux.Vars(r)["nombre"]

    c := models.Course{Nombre: nombre}
    if err := c.GetCourse(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            utils.RespondWithError(w, http.StatusNotFound, "Course not found")
        default:
            utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) GetCourses(w http.ResponseWriter, r *http.Request) {

    courses, err :=  models.GetCoursesU(a.DB)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, courses)
}



func (a *App) UpdateCourse(w http.ResponseWriter, r *http.Request) {
    var c models.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
    err := c.UpdateCourse(a.DB);
    if  err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) DeleteCourse(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	nombre := vars["nombre"]

    p := models.Course{Nombre: nombre}
    if err := p.DeleteCourse(a.DB); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}