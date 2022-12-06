package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Akanom/Golang-Auto-Shop-Management-APIs/pkg/models"
	"github.com/Akanom/Golang-Auto-Shop-Management-APIs/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewAuto models.Auto

func GetAuto(w http.ResponseWriter, r *http.Request) {
	newAutos := models.GetAuto()
	res, _ := json.Marshal(newAutos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAutoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	autoId := vars["autoId"]
	ID, err := strconv.ParseInt(autoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	autoDetails, _ := models.GetAutoById(ID)
	res, _ := json.Marshal(autoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNewAuto(w http.ResponseWriter, r *http.Request) {
	CreateNewAuto := &models.Auto{}
	utils.ParseBody(r, CreateNewAuto)
	b := CreateNewAuto.CreateNewAuto()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	autoId := vars["autoId"]
	ID, err := strconv.ParseInt(autoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	auto := models.DeleteAuto(ID)
	res, _ := json.Marshal(auto)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAuto(w http.ResponseWriter, r *http.Request) {
	var updateAuto = &models.Auto{}
	utils.ParseBody(r, UpdateAuto)
	vars := mux.Vars(r)
	autoId := vars["autoId"]
	ID, err := strconv.ParseInt(autoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	autoDetails, db := models.GetAutoById(ID)
	if updateAuto.Name != "" {
		autoDetails.Name = updateAuto.Name
	}
	if updateAuto.AutoModel != "" {
		autoDetails.AutoModel = updateAuto.AutoModel
	}
	if updateAuto.ManufactureDate != "" {
		autoDetails.ManufactureDate = updateAuto.ManufactureDate
	}
	db.Save(&autoDetails)
	res, _ := json.Marshal(autoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
