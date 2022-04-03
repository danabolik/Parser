package repositories

import "Parser/internal/domain"

func Insert(file domain.File) bool {
	file.GetFileName()
	file.GetExtension()
	//TODO: Доработать репозиторий
	return true
}
