package domain

import (
	"mime/multipart"
)

// ICsvService définit l'interface pour le service de traitement des fichiers CSV.
type ICsvService interface {
	ProcessCsv(file multipart.File) error
}
