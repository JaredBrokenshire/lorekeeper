package file_storage

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileContext interface {
	File(file string) error
}

var (
	ErrFileRead   = errors.New("error reading file")
	ErrFileDelete = errors.New("error deleting file")
	ErrDirDelete  = errors.New("error deleting directory")
	ErrFileSave   = errors.New("error saving file")
)

type LocalStorage struct {
	Storage string
}

var folderLocations = []string{"../", "../../"}

func NewLocalStorage(localStoragePath string) *LocalStorage {

	// Let's find the main folder
	locationParts := strings.Split(localStoragePath, "/")

	if _, err := os.Stat(locationParts[0]); os.IsNotExist(err) {
		// Loop through spossible folder locations checking for the folder
		for _, folderLocation := range folderLocations {
			if _, err := os.Stat(folderLocation + locationParts[0]); !os.IsNotExist(err) {
				// If this folder exists then update the localStoragePath
				localStoragePath = folderLocation + localStoragePath
				log.Printf("Found location: updating location storage path to %v\n", localStoragePath)
				break
			}
		}
	}

	// make sure the storage directory exists
	if _, err := os.Stat(localStoragePath); os.IsNotExist(err) {
		err := os.Mkdir(localStoragePath, 0755)
		if err != nil {
			log.Fatalf("FATAL: Error starting local storage service. Unable to create folder for local handler: %v", err)
		}
	}
	return &LocalStorage{
		Storage: localStoragePath,
	}
}

func (s *LocalStorage) Get(c FileContext, path string) error {
	return c.File(s.Storage + path)
}

func (s *LocalStorage) ReadFile(path string) ([]byte, error) {
	fullPath := filepath.Join(s.Storage, path)
	res, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, ErrFileRead
	}
	return res, nil
}

func (s *LocalStorage) Save(path string, file io.Reader, filename string) error {
	uploadPath := filepath.Join(s.Storage, path)
	err := os.MkdirAll(uploadPath, 0755)
	if err != nil {
		return ErrFileSave
	}

	// Destination
	dst, err := os.Create(filepath.Join(uploadPath, filename))
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return ErrFileSave
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, file); err != nil {
		return ErrFileSave
	}

	return nil
}

func (s *LocalStorage) Delete(path string) error {
	err := os.Remove(filepath.Join(s.Storage, path))
	if err != nil {
		return ErrFileDelete
	}
	return nil
}

func (s *LocalStorage) DeleteDir(path string) error {
	err := os.RemoveAll(filepath.Join(s.Storage, path))
	if err != nil {
		return ErrDirDelete
	}
	return nil
}

func (s *LocalStorage) test() error {
	return nil
}
