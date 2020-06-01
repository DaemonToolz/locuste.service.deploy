package main

import (
	"fmt"
	"os"
)

// StartInstallProcedure Démarrer la procédure d'installation
func StartInstallProcedure(version string) {
	origin := fmt.Sprintf("./repo/versions/%s/", version)
	destination := "/home/pi/Documents/locuste/"
	status := ProgressIndicator{
		Status:  InProgress,
		Message: "Suppression de l'ancienne version",
	}
	broadcastUpdate("install", status)
	RemoveContents(destination)
	os.Remove(destination)

	status.Message = "Préparation de l'installation"
	broadcastUpdate("install", status)

	indicator := &FileCopyInfo{
		FileIndex: 0,
		FileCount: CountFiles(origin, origin),
	}

	status.Message = "Installation en cours"
	broadcastUpdate("install", status)
	err := CopyDirectory(origin, destination, indicator)
	if err != nil {
		status.Status = Error
		status.Message = "Une erreur s'est produite durant l'installation"
	} else {
		status.Status = Success
		status.Message = "Installation terminée avec succès"
	}
	broadcastUpdate("install", status)
}
