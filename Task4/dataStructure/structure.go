package structure

import (
	"strconv"
	"task4/database"
	filesystem "task4/fileSystem"
	"task4/models"
)

type FileStructure struct {
}

type DatabaseStructure struct {
}

type Operations interface {
	GetData() error
	PutData(data models.Data) error
}

func (fs FileStructure) GetData() error {
	return filesystem.FetchData()
}

func (fs FileStructure) PutData(data models.Data) error {
	filedata := strconv.Itoa(data.UserId) + "\t" + strconv.Itoa(data.Id) + "\t" + data.Title + "\n"
	return filesystem.WriteOnFile(filedata)
}

func (db DatabaseStructure) GetData() error {
	return database.FindData()
}

func (db DatabaseStructure) PutData(data models.Data) error {
	return database.InsertData(data)
}
