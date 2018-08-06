package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

type JobAd struct {
	Id    int    `json:"id"`
	title string `json:"title"`
}

var jobads []JobAd

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/jobad", GetJobAds).Methods("GET")
	router.HandleFunc("/jobad/{id}", GetJobAd).Methods("GET")
	router.HandleFunc("/jobad/", CreateJobAd).Methods("POST")
	router.HandleFunc("/jobad/{id}", DeleteJobAd).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetJobAds(w http.ResponseWriter, r *http.Request) {
	jobads = append(jobads, JobAd{Id: 1, title: "Desenvolvedor Gopher"})
	json.NewEncoder(w).Encode(jobads)
}

func GetJobAd(w http.ResponseWriter, r *http.Request) {
	jobads = append(jobads, JobAd{Id: 1, title: "Desenvolvedor Gopher"})
	json.NewEncoder(w).Encode(jobads)
}

func CreateJobAd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	errors := validatePost(string(body))

	if len(errors) > 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(errors)
	} else {
		var jobAd JobAd
		err = json.Unmarshal(body, &jobAd)
		fmt.Println(jobAd)

		if err != nil {
			panic(err)
		}
		jobAd.Id = len(jobads) + 1
		jobads = append(jobads, jobAd)
		json.NewEncoder(w).Encode(jobads)
	}
}

func validatePost(post string) []string {
	pwd, _ := os.Getwd()
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + pwd + "/schema/schema.json")
	stringLoader := gojsonschema.NewStringLoader(post)

	result, err := gojsonschema.Validate(schemaLoader, stringLoader)
	if err != nil {
		panic(err.Error())
	}
	var errors []string
	if result.Valid() {
		return errors
	}

	for _, desc := range result.Errors() {
		errors = append(errors, desc.Description())
	}
	return errors
}

func DeleteJobAd(w http.ResponseWriter, r *http.Request) {}
