package service

import (
	"context"
	"fmt"
	"strings"

	"projek-backend/app/model"
	"projek-backend/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Register(ctx context.Context, request *model.RegisterRequest) (*model.User, error) {
	if err := ValidateRegister(request.Username, request.Password); err != nil {
		return nil, err
	}

	username := strings.TrimSpace(request.Username)

	_, err := s.userRepository.FindByUsername(ctx, username)
	if err == nil {
		return nil, fmt.Errorf("username sudah digunakan")
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan hashing password: %w", err)
	}

	role := "customer"

	user := &model.User{
		Username: username,
		Password: string(passwordHash),
		Role:     role,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, request *model.LoginRequest) (*model.User, error) {
	if err := ValidateLogin(request.Username, request.Password); err != nil {
		return nil, err
	}

	username := strings.TrimSpace(request.Username)

	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	); err != nil {
		return nil, fmt.Errorf("username atau password salah")
	}

	user.Password = ""

	return user, nil
}