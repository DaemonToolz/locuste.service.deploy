package main

// FileCopyInfo Informations sur l'état d'avancement des copies de fichier
type FileCopyInfo struct {
	CurrentFile string `json:"current_file"`
	FileCount   int    `json:"file_count"`
	FileIndex   int    `json:"fil_index"`
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
