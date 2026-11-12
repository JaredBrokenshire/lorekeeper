package mocks

import (
	"lorekeeper/db/models"
	"lorekeeper/services"
)

type MockFileService struct {
	CreateCalls []services.CreateFileDTO
	CreateError error

	ValidateCalls []struct {
		File         services.UploadedFileDTO
		AllowedTypes []string
	}
	ValidateError error
}

func NewMockFileService() *MockFileService {
	res := &MockFileService{}
	res.Reset()
	return res
}

func (m *MockFileService) Create(dto services.CreateFileDTO) (*models.File, error) {
	m.CreateCalls = append(m.CreateCalls, dto)

	return nil, m.CreateError
}

func (m *MockFileService) Validate(file services.UploadedFileDTO, allowedTypes []string) error {
	m.ValidateCalls = append(m.ValidateCalls, struct {
		File         services.UploadedFileDTO
		AllowedTypes []string
	}{file, allowedTypes})

	return m.ValidateError
}

func (m *MockFileService) Reset() {
	m.CreateCalls = []services.CreateFileDTO{}
	m.CreateError = nil

	m.ValidateCalls = []struct {
		File         services.UploadedFileDTO
		AllowedTypes []string
	}{}
	m.ValidateError = nil
}
