package service

import (
	domain "api/internal/domain/services"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
)

// CsvService est l'implémentation concrète de ICsvService.
type CsvService struct{}

// NewCsvService crée une nouvelle instance de CsvService.
func NewCsvService() domain.ICsvService {
	return &CsvService{}
}

// ProcessCsv traite le fichier CSV téléchargé.
func (cs *CsvService) ProcessCsv(file multipart.File) error {
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
	return nil
}
