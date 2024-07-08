package auth

import (
	"auth_service/internal/domain"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Auth struct {
	log *slog.Logger
	s   Storage
}

type Storage interface {
	CreateUser(user *domain.Users) (int64, error)
	FindUser(user *domain.Users) (int64, string, bool, error)
}

func New(log *slog.Logger, s Storage) *Auth {
	return &Auth{
		log: log,
		s:   s,
	}
}

func (a *Auth) Register(ctx context.Context, name, surname, login, password, mail string) (int64, error) {
	a.log.Info("registering user")

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.log.Error("could not hash password")
		return 0, fmt.Errorf("could not hash password: %w", err)
	}
	user := &domain.Users{
		Name:     name,
		Surname:  surname,
		Login:    login,
		Password: string(hashPass),
		Email:    mail,
	}

	a.log.Info("creating user", "user", user)
	id, err := a.s.CreateUser(user)
	if err != nil {
		a.log.Error("failed to create user")
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (a *Auth) Login(ctx context.Context, login, password string) (int64, bool, error) {
	a.log.Info("login user", "login", login)

	id, hashPass, enable, err := a.s.FindUser(&domain.Users{Login: login})
	if err != nil {
		a.log.Error("failed to find user")
		return 0, false, fmt.Errorf("failed to find user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		a.log.Error("invalid password")
		return 0, false, fmt.Errorf("invalid password: %w", err)
	}
	return id, enable, nil
}

func (a *Auth) Generate(ctx context.Context, login string) (string, error) {
	a.log.Info("generating code", "login", login)
	panic("implement me")
}

func (a *Auth) VerifyQr(ctx context.Context, username, code string) (int64, string, error) {
	panic("implement me")
}
