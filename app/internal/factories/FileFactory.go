package factories

import (
	"Parser/internal/domain"
	"errors"
)

func CreateByExtension(ext string) (domain.File, error) {
	switch ext {
	case ".csv":
		return &domain.CsvFile{}, nil
	default:
		return nil, errors.New("file extension can not parse")
	}
}
