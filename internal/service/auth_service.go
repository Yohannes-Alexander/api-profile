package service

import (
	"errors"
	"github.com/Yohannes-Alexander/api-profile/internal/dto"
	"github.com/Yohannes-Alexander/api-profile/internal/repository"		
	"github.com/Yohannes-Alexander/api-profile/internal/utils"	
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.TokenResponse, error)
	RefreshToken(refreshToken string) (*dto.TokenResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{repo: r}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	accessToken, refreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}
	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*dto.TokenResponse, error) {
	userID, err := utils.ValidateToken(refreshToken, true)
	if err != nil {
		return nil, err
	}
	accessToken, newRefreshToken, err := utils.GenerateTokens(userID)
	if err != nil {
		return nil, err
	}
	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
