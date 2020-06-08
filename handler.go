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
	fnameArr := strings.Split(header.Filename, ".")
	go Unzip(sourcePath, fmt.Sprintf("./repo/versions/%s/%s", vars["version"], strings.Join(fnameArr[:len(fnameArr)-1], ".")))
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

// GetInstalledVersion Récupère la version installée
func GetInstalledVersion(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(GetDiskVersion(false)); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}

// Uninstall Débute la procédure de désinstallation
func Uninstall(w http.ResponseWriter, r *http.Request) {

	go StartUninstallProcedure()

	if err := json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
	}{true}); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}

// RunCommand Démarre ou arrête un processus (exécute une commande)
func RunCommand(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var post ExecCommand
	err := decoder.Decode(&post)

	if err != nil {
		if err := json.NewEncoder(w).Encode(struct {
			Success bool `json:"failed"`
		}{true}); err != nil {
			failOnError(err, "Unable to load the message")
			panic(err)
		}
	}

	go Run(post)

	if err := json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
	}{true}); err != nil {
		failOnError(err, "Unable to load the message")
		panic(err)
	}
}
