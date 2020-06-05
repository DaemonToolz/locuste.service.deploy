package main

// FileCopyInfo Informations sur l'état d'avancement des copies de fichier
type FileCopyInfo struct {
	CurrentFile string `json:"current_file"`
	FileCount   int    `json:"file_count"`
	FileIndex   int    `json:"file_index"`
}

// OperationStatus Statut de l'opération
type OperationStatus int

const (
	// Success Succès
	Success OperationStatus = iota
	// InProgress En cours
	InProgress OperationStatus = iota
	// Error En erreur
	Error OperationStatus = iota
)

// ProgressIndicator Indicateur d'avancement sur une opération
type ProgressIndicator struct {
	Message string          `json:"message"`
	Status  OperationStatus `json:"status"`
}

// ProjectVersion Version globale du projet
type ProjectVersion struct {
	GlobalVersion   string       `json:"global_version"`
	DetailedVersion []AppVersion `json:"detailed_version"`
}

// AppVersion Version globale du projet
type AppVersion struct {
	Name    string `json:"app_name"`
	Version string `json:"version"`
}
