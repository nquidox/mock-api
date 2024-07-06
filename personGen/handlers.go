package personGen

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()
	log.Info("Query Params: ", queryParams)

	passportSerie := queryParams.Get("passportSerie")
	ps, err := strconv.Atoi(passportSerie)
	if err != nil {
		w.WriteHeader(400)
	}

	passportNumber := queryParams.Get("passportNumber")
	pn, err := strconv.Atoi(passportNumber)
	if err != nil {
		w.WriteHeader(400)
	}

	log.Debug("Request: ", r.URL.Path)
	log.Debug("Query Parameters: ", r.URL.Query())

	person := Person{}
	person.PassportSerie = ps
	person.PassportNumber = pn

	err = person.Read()
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
	}

	if person.Name == "" {
		person.Create()
		person.Save()
	}

	log.WithFields(log.Fields{
		"ФИО":     fmt.Sprintf("%s %s %s", person.Name, person.Patronymic, person.Surname),
		"Город":   person.Address,
		"Паспорт": fmt.Sprintf("%d %d", person.PassportSerie, person.PassportNumber),
	}).Info("JSON Response")
	bytes, err := json.Marshal(&person)
	if err != nil {
		log.Error(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Error(err)
	}
}
