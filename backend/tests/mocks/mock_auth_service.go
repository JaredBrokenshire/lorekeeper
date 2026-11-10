package mocks

import (
	"lorekeeper/db/models"
	"lorekeeper/services"
)

type MockAuthService struct {
	RegisterCalls []services.RegisterRequestDTO
	RegisterError error

	LoginCalls []services.LoginRequestDTO
	LoginError error

	RefreshCalls []string
	RefreshError error

	GenerateAccessTokenCalls []*models.User
	GenerateAccessTokenError error

	GenerateRefreshTokenCalls []*models.User
	GenerateRefreshTokenError error
}

func NewMockAuthService() *MockAuthService {
	return &MockAuthService{
		RegisterCalls: []services.RegisterRequestDTO{},
		RegisterError: nil,

		LoginCalls: []services.LoginRequestDTO{},
		LoginError: nil,

		RefreshCalls: []string{},
		RefreshError: nil,

		GenerateAccessTokenCalls: []*models.User{},
		GenerateAccessTokenError: nil,

		GenerateRefreshTokenCalls: []*models.User{},
		GenerateRefreshTokenError: nil,
	}
}

func (m *MockAuthService) Register(dto services.RegisterRequestDTO) (*services.AuthResponse, string, error) {
	m.RegisterCalls = append(m.RegisterCalls, dto)
	return nil, "", m.RegisterError
}

func (m *MockAuthService) Login(dto services.LoginRequestDTO) (*services.AuthResponse, string, error) {
	m.LoginCalls = append(m.LoginCalls, dto)
	return nil, "", m.LoginError
}

func (m *MockAuthService) Refresh(refreshToken string) (*services.AuthResponse, string, error) {
	m.RefreshCalls = append(m.RefreshCalls, refreshToken)
	return nil, "", m.RefreshError
}

func (m *MockAuthService) GenerateAccessToken(user *models.User) (string, error) {
	m.GenerateAccessTokenCalls = append(m.GenerateAccessTokenCalls, user)
	return "", m.GenerateAccessTokenError
}

func (m *MockAuthService) GenerateRefreshToken(user *models.User) (string, error) {
	m.GenerateRefreshTokenCalls = append(m.GenerateRefreshTokenCalls, user)
	return "", m.GenerateRefreshTokenError
}

func (m *MockAuthService) Reset() {
	m.RegisterCalls = []services.RegisterRequestDTO{}
	m.RegisterError = nil

	m.LoginCalls = []services.LoginRequestDTO{}
	m.LoginError = nil

	m.RefreshCalls = []string{}
	m.RefreshError = nil

	m.GenerateAccessTokenCalls = []*models.User{}
	m.GenerateAccessTokenError = nil

	m.GenerateRefreshTokenCalls = []*models.User{}
	m.GenerateRefreshTokenError = nil
}
