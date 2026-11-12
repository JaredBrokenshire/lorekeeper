package mocks

import (
	"io"
	"lorekeeper/pkg/file_storage"
)

type MockFileStore struct {
	GetCalls []string
	GetError error

	ReadFileCalls    []string
	ReadFileResponse []byte
	ReadFileError    error

	SaveCalls []struct{ Path, FileName string }
	SaveError error

	DeleteCalls []string
	DeleteError error

	DeleteDirCalls []string
	DeleteDirError error
}

func NewMockFileStore() *MockFileStore {
	m := MockFileStore{}
	m.Reset()
	return &m
}

func (m *MockFileStore) Get(_ file_storage.FileContext, file string) error {
	m.GetCalls = append(m.GetCalls, file)
	return m.GetError
}

func (m *MockFileStore) ReadFile(path string) ([]byte, error) {
	m.ReadFileCalls = append(m.ReadFileCalls, path)
	return m.ReadFileResponse, m.ReadFileError
}

func (m *MockFileStore) Save(path string, _ io.Reader, fileName string) error {
	m.SaveCalls = append(m.SaveCalls, struct{ Path, FileName string }{path, fileName})
	return m.SaveError
}

func (m *MockFileStore) Delete(path string) error {
	m.DeleteCalls = append(m.DeleteCalls, path)
	return m.DeleteError
}

func (m *MockFileStore) DeleteDir(path string) error {
	m.DeleteDirCalls = append(m.DeleteDirCalls, path)
	return m.DeleteDirError
}

func (m *MockFileStore) Reset() {
	m.GetCalls = []string{}
	m.GetError = nil

	m.ReadFileCalls = []string{}
	m.ReadFileResponse = []byte("A file used for testing")
	m.ReadFileError = nil

	m.SaveCalls = []struct{ Path, FileName string }{}
	m.SaveError = nil

	m.DeleteCalls = []string{}
	m.DeleteError = nil

	m.DeleteDirCalls = []string{}
	m.DeleteDirError = nil
}
