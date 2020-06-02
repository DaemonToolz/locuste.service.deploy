package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Upload Charge une nouvelle version de l'exécutable
func Upload(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	r.ParseMultipartForm(32 << 15) // limit your max input length!

	file, header, err := r.FormFile("file")
	if file == nil || err != nil {
		if err := json.NewEncoder(w).Encode(struct {
			Success bool `json:"failed"`
		}{true}); err != nil {
			failOnError(err, "Unable to load the message")
			panic(err)
		}
		return
	}
	defer file.Close()
	sourcePath := fmt.Sprintf("./repo/archive/%s_%s", vars["version"], header.Filename)
	out, err := os.Create(sourcePath)
	defer out.Close()

	if err != nil {
		if err := json.NewEncoder(w).Encode(struct {
			Success bool `json:"failed"`
		}{true}); err != nil {
			failOnError(err, "Unable to load the message")
			panic(err)
		}
		return
	}

	io.Copy(out, file)
	go Unzip(sourcePath, fmt.Sprintf("./repo/versions/%s/%s", vars["version"], strings.Split(header.Filename, ".")[0]))
	if err := json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
	}{true}); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}

// Install Débute la procédure d'installation
func Install(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	go StartInstallProcedure(vars["version"])

	if err := json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
	}{true}); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}

// DeleteVersion Supprimer une version installée
func DeleteVersion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	go DeleteVersionFiles("./repo/archive", vars["version"])
	go DeleteVersionFiles("./repo/versions", vars["version"])

	if err := json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
	}{true}); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}

// GetAvailableVersions Débute la procédure d'installation
func GetAvailableVersions(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(ListVersions()); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}
