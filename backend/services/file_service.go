package services

import (
	"errors"
	"io"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"slices"
)

type FileServiceInterface interface {
	Create(file CreateFileDTO) (*models.File, error)
	Validate(file UploadedFileDTO, allowedTypes []string) error
}

var (
	ErrFileTooLarge    = errors.New("file too large")
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileRepo        = errors.New("file repository error")
)

type FileService struct {
	fileRepo    *repositories.FileRepository
	userRepo    *repositories.UserRepository
	MaxFileSize int64
}

func NewFileService(fileRepo *repositories.FileRepository, userRepo *repositories.UserRepository) *FileService {
	return &FileService{
		fileRepo: fileRepo,
		userRepo: userRepo,

		MaxFileSize: 10 * 1024 * 1024, // 10MB
	}
}

type UploadedFileDTO struct {
	Name        string
	ContentType string
	Size        int64
	Open        func() (io.ReadCloser, error)
}

type CreateFileDTO struct {
	Name     string
	Location string
	AuthorID uint
}

func (s *FileService) Create(dto CreateFileDTO) (*models.File, error) {
	user := s.userRepo.GetById(dto.AuthorID)
	if user == nil {
		return nil, ErrUserNotFound
	}

	fileModel := models.File{
		Name:     dto.Name,
		Location: dto.Location,
		Author:   user,
	}

	err := s.fileRepo.Create(&fileModel)
	if err != nil {
		return nil, ErrFileRepo
	}

	return &fileModel, nil
}

func (s *FileService) Validate(file UploadedFileDTO, allowedTypes []string) error {
	if file.Size > s.MaxFileSize {
		return ErrFileTooLarge
	}

	if !slices.Contains(allowedTypes, file.ContentType) {
		return ErrInvalidFileType
	}

	return nil
}
