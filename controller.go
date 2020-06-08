package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// StartInstallProcedure Démarrer la procédure d'installation
func StartInstallProcedure(version string) {
	origin := "./repo/versions/"
	destination := "/home/pi/Documents/locuste/"
	StartUninstallProcedure()

	status := ProgressIndicator{
		Status:  InProgress,
		Message: "Préparation de l'installation",
	}
	broadcastUpdate("install", status)

	myVersion := fmt.Sprintf("%s%s", origin, version)
	indicator := &FileCopyInfo{
		FileIndex: 0,
		FileCount: CountFiles(myVersion, myVersion),
	}

	status.Message = "Installation en cours"
	broadcastUpdate("install", status)
	err := CopyDirectory(origin, version, destination, true, indicator)
	if err != nil {
		status.Status = Error
		status.Message = "Une erreur s'est produite durant l'installation"
	} else {
		status.Status = Success
		status.Message = "Installation terminée avec succès"
	}
	broadcastUpdate("install", status)
}

// StartUninstallProcedure Démarrer la procédure de désinstallation
func StartUninstallProcedure() {
	destination := "/home/pi/Documents/locuste/"
	status := ProgressIndicator{
		Status:  InProgress,
		Message: "Désinstallation de la version actuelle",
	}
	broadcastUpdate("install", status)

	RemoveContents(destination)
	os.Remove(destination)

	status.Status = Success
	status.Message = "Désinstallation terminée avec succès"

	broadcastUpdate("install", status)
}

// GetDiskVersion Récupère la version installée sur le disque
func GetDiskVersion(rootVersion bool) ProjectVersion {

	entries, err := ioutil.ReadDir("/home/pi/Documents/locuste/")
	// On ne devrait avoir qu'un seul résultat

	version := ProjectVersion{
		GlobalVersion: "N/A",
	}

	if err != nil {
		return version
	}

	version.DetailedVersion = make([]AppVersion, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := fmt.Sprintf("/home/pi/Documents/locuste/%s", entry.Name())

		version.GlobalVersion = entry.Name()
		if !rootVersion {
			extractApps(path, &version)
		}

	}
	return version
}

func extractApps(path string, pVersion *ProjectVersion) {

	apps, err := ioutil.ReadDir(path)
	if err != nil {
		failOnError(err, "Une erreur est survenue lors de l'extraction de l'application")
		return
	}

	for _, app := range apps {
		if !app.IsDir() {
			continue
		}
		extractAppVersion(path, app.Name(), pVersion)
	}
}

func extractAppVersion(path, app string, pVersion *ProjectVersion) {
	appPath := fmt.Sprintf("%s/%s", path, app)
	appVersions, err := ioutil.ReadDir(appPath)
	if err != nil {
		failOnError(err, "Une erreur est survenue lors de l'extraction de la version")
		return
	}
	for _, appVersion := range appVersions {
		if !appVersion.IsDir() {
			continue
		}

		proc := getProcess(app)
		var path string
		var err error
		if proc != nil {
			path, err = proc.Path()
		}

		pVersion.DetailedVersion = append(pVersion.DetailedVersion, AppVersion{
			Name:      app,
			IsRunning: proc != nil && err == nil && proc.Pid() > 0 && strings.Contains(path, appPath),
			Version:   appVersion.Name(),
		})
	}

}
