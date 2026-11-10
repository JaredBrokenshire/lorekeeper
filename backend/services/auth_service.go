package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/middleware"
	"strconv"
	"time"
)

type AuthServiceInterface interface {
	Register(dto RegisterRequestDTO) (*AuthResponse, string, error)
	Login(dto LoginRequestDTO) (*AuthResponse, string, error)
	Refresh(refreshToken string) (*AuthResponse, string, error)
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
}

type AuthResponse struct {
	AccessToken string
	User        *models.User
}

var (
	ErrUsernameTaken      = errors.New("username already taken")
	ErrEmailTaken         = errors.New("email already registered")
	ErrPasswordNotMatch   = errors.New("passwords do not match")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthService struct {
	db        *gorm.DB
	userRepo  *repositories.UserRepository
	roleRepo  *repositories.RoleRepository
	jwtSecret []byte
}

func NewAuthService(db *gorm.DB, userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

type RegisterRequestDTO struct {
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

func (s *AuthService) Register(dto RegisterRequestDTO) (*AuthResponse, string, error) {
	if dto.Password != dto.ConfirmPassword {
		return nil, "", ErrPasswordNotMatch
	}

	existingUser := s.userRepo.GetByUsername(dto.Username)
	if existingUser.ID != 0 {
		return nil, "", ErrUsernameTaken
	}

	existingEmail := s.userRepo.GetByEmail(dto.Email)
	if existingEmail.ID != 0 {
		return nil, "", ErrEmailTaken
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	role := s.roleRepo.GetByName(models.RoleUser)

	user := &models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: string(passwordHash),
		Roles:    []*models.Role{role},
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, refreshToken, nil
}

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthService) Login(dto LoginRequestDTO) (*AuthResponse, string, error) {
	user := s.userRepo.GetByUsername(dto.Username)
	if user.ID == 0 {
		return nil, "", ErrInvalidCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, refreshToken, nil
}

func (s *AuthService) Refresh(refreshToken string) (*AuthResponse, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, "", ErrInvalidCredentials
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, "", ErrInvalidCredentials
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, "", ErrInvalidCredentials
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	user := s.userRepo.GetById(userId)
	if user.ID == 0 {
		return nil, "", ErrUserNotFound
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, refreshToken, nil
}

func (s *AuthService) GenerateAccessToken(user *models.User) (string, error) {
	claims := middleware.JWTClaims{
		UserID: user.ID,

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "localdnd-api",
			Subject:  strconv.FormatUint(uint64(user.ID), 10),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(15 * time.Minute),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) GenerateRefreshToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  strconv.FormatUint(uint64(user.ID), 10),
		"type": "refresh",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iss":  "localdnd-api",
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(s.jwtSecret)
}
