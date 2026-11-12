package services

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/file_storage"
	"os"
)

type Dependencies struct {
	db        *gorm.DB
	repos     *repositories.Repos
	jwtSecret string

	// File Storage
	fileStore file_storage.Store

	// Services
	authService      AuthServiceInterface
	userService      UserServiceInterface
	characterService CharacterServiceInterface
	campaignService  CampaignServiceInterface
	fileService      FileServiceInterface
	documentService  DocumentServiceInterface
}

func NewDependencies(db *gorm.DB, repos *repositories.Repos) *Dependencies {
	return &Dependencies{
		db:        db,
		repos:     repos,
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (d *Dependencies) GetDB() *gorm.DB {
	return d.db
}

func (d *Dependencies) GetFileStore() file_storage.Store {
	if d.fileStore == nil {
		d.fileStore = file_storage.NewLocalStorage("assets/files")
	}

	return d.fileStore
}

func (d *Dependencies) SetFileStore(fileStore file_storage.Store) {
	d.fileStore = fileStore
}

func (d *Dependencies) GetAuthService() AuthServiceInterface {
	if d.authService == nil {
		d.authService = NewAuthService(d.db, d.repos.User, d.repos.Role, d.jwtSecret)
	}

	return d.authService
}

func (d *Dependencies) SetAuthService(authService AuthServiceInterface) {
	d.authService = authService
}

func (d *Dependencies) GetUserService() UserServiceInterface {
	if d.userService == nil {
		d.userService = NewUserService(d.repos.User)
	}

	return d.userService
}

func (d *Dependencies) SetUserService(userService UserServiceInterface) {
	d.userService = userService
}

func (d *Dependencies) GetCharacterService() CharacterServiceInterface {
	if d.characterService == nil {
		d.characterService = NewCharacterService(d.repos.Character)
	}

	return d.characterService
}

func (d *Dependencies) SetCharacterService(characterService CharacterServiceInterface) {
	d.characterService = characterService
}

func (d *Dependencies) GetCampaignService() CampaignServiceInterface {
	if d.campaignService == nil {
		d.campaignService = NewCampaignService(
			d.repos.Campaign,
			d.GetFileService(),
			d.GetFileStore(),
		)
	}

	return d.campaignService
}

func (d *Dependencies) SetCampaignService(campaignService CampaignServiceInterface) {
	d.campaignService = campaignService
}

func (d *Dependencies) GetFileService() FileServiceInterface {
	if d.fileService == nil {
		d.fileService = NewFileService(d.repos.File, d.repos.User)
	}

	return d.fileService
}

func (d *Dependencies) SetFileService(fileService FileServiceInterface) {
	d.fileService = fileService
}

func (d *Dependencies) GetDocumentService() DocumentServiceInterface {
	if d.documentService == nil {
		d.documentService = NewDocumentService()
	}

	return d.documentService
}

func (d *Dependencies) SetDocumentService(documentService DocumentServiceInterface) {
	d.documentService = documentService
}
