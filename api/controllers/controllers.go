package controllers

import(
  "net/http"
  "api/models"
  "encoding/json"
  "github.com/gorilla/mux"
  "api/database"
)

func AllPersonas(res http.ResponseWriter, req *http.Request){
  var p []models.Personalites

  database.DB.Find(&p)
  json.NewEncoder(res).Encode(p)
}

func OnePersona(res http.ResponseWriter, req *http.Request){
  vars := mux.Vars(req)
  id := vars["id"]

  var p models.Personalites
  database.DB.First(&p, id)
  json.NewEncoder(res).Encode(p)
}

func Register(res http.ResponseWriter, req *http.Request){
  var newPersona models.Personalites

  json.NewDecoder(req.Body).Decode(&newPersona)
  database.DB.Create(&newPersona)
  json.NewEncoder(res).Encode(newPersona)
}

func Delete(res http.ResponseWriter, req *http.Request){
  vars := mux.Vars(req)
  id := vars["id"]

  var deletedPersona models.Personalites
  database.DB.Delete(&deletedPersona, id)
  json.NewEncoder(res).Encode(deletedPersona)
}

func Edit(res http.ResponseWriter, req *http.Request)  {

  vars := mux.Vars(req)
  id := vars["id"]

  var editPersona models.Personalites

  database.DB.First(&editPersona, id)
  json.NewDecoder(req.Body).Decode(&editPersona)
  database.DB.Save(&editPersona)

  json.NewEncoder(res).Encode(editPersona)
}
