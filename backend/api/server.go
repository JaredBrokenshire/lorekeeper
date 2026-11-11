package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	database "lorekeeper/db"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/file_storage"
	"lorekeeper/services"
	"time"
)

type Server struct {
	Echo         *echo.Echo
	DB           *gorm.DB
	Repos        *repositories.Repos
	Dependencies *services.Dependencies
}

func NewServer() *Server {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	fileStore := file_storage.NewLocalStorage("assets/files")

	db, err := database.Init(fileStore)
	if err != nil {
		panic(err)
	}

	s := &Server{
		Echo: echo.New(),
		DB:   db,
	}

	s.Repos = repositories.NewRepos(s.DB)
	s.Dependencies = services.NewDependencies(s.DB, s.Repos)
	s.Dependencies.SetFileStore(fileStore)

	return s
}

func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
