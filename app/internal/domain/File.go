package domain

type File interface {
	GetId() uint
	GetFileName() string
	SetFileName(fileName string)
	GetExtension() string
	SetExtension(ext string)
}

func (f CsvFile) GetId() uint {
	return f.Id
}

func (f CsvFile) GetFileName() string {
	return f.fileName
}
func (f *CsvFile) SetFileName(fileName string) {
	f.fileName = fileName
}

func (f CsvFile) GetExtension() string {
	return f.extension
}
func (f *CsvFile) SetExtension(ext string) {
	f.extension = ext
}

type CsvFile struct {
	Id        uint
	fileName  string
	extension string
}

type FileRow struct {
	FileId uint
	Data   string
}
