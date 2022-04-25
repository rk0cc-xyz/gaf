package storage

import (
	"time"

	"github.com/rk0cc-xyz/gaf/structure"
)

// AA structure contains raw data from database.
type DatabaseFieldContainer struct {
	page      int64
	content   []byte
	updatedAt string
}

// An interface for handling read and write data between the container and database.
type DataBaseFieldHandler interface {
	// Handle how the content saved into database.
	WriteToDB(page int64, content []byte, updatedAt string) error

	// Handle how to parse data from database.
	ReadFromDB(page int64) ([]byte, *string, error)
}

// Construct a field container by providing page and content.
func CreateDatabaseFieldContainer(page int64, content []structure.GitHubRepositoryStructure) (*DatabaseFieldContainer, error) {
	compressedContent, ccerr := compressContent(content)
	if ccerr != nil {
		return nil, ccerr
	}

	return &DatabaseFieldContainer{
		page:      page,
		content:   compressedContent,
		updatedAt: time.Now().UTC().Format(time.RFC3339),
	}, nil
}

// Get the field data from database with the handler.
func GetFieldContainerFromDatabase(page int64, handler DataBaseFieldHandler) (*DatabaseFieldContainer, error) {
	c, u, err := handler.ReadFromDB(page)
	if err != nil {
		return nil, err
	}

	return &DatabaseFieldContainer{
		page:      page,
		content:   c,
		updatedAt: *u,
	}, nil
}

// Get context page from API.
func (dfc DatabaseFieldContainer) GetPage() int64 {
	return dfc.page
}

// Content of the API result.
func (dfc DatabaseFieldContainer) GetContent() ([]structure.GitHubRepositoryStructure, error) {
	return decompressContent(dfc.content)
}

// Timestamp when this container created.
func (dfc DatabaseFieldContainer) GetUpdatedAt() string {
	return dfc.updatedAt
}

// Archive container's data to database.
func (dfc DatabaseFieldContainer) SaveToDatabase(handler DataBaseFieldHandler) error {
	return handler.WriteToDB(dfc.page, dfc.content, dfc.updatedAt)
}
