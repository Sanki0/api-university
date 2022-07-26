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

func (a *App) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var c models.Record
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
	if err := c.CreateRecord(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, c)
}

func (a *App) GetRecord(w http.ResponseWriter, r *http.Request) {
    student := mux.Vars(r)["student"]
	course := mux.Vars(r)["course"]


    c := models.Record{Student: student, Course: course}
    if err := c.GetRecord(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            utils.RespondWithError(w, http.StatusNotFound, "Record not found")
        default:
            utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) GetRecords(w http.ResponseWriter, r *http.Request) {

    records, err :=  models.GetRecordsU(a.DB)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, records)
}



func (a *App) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var c models.Record
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
    
    err := c.UpdateRecord(a.DB);
    if  err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    student := mux.Vars(r)["student"]
	course := mux.Vars(r)["course"]

    p := models.Record{Student: student, Course: course}
    if err := p.DeleteRecord(a.DB); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}