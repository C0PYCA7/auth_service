package postgres

import (
	"auth_service/internal/config"
	"auth_service/internal/domain"
	"auth_service/internal/storage"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(dbConfig config.Database) (*Storage, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database,
	)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(user *domain.Users) (int64, error) {
	result := s.db.Create(user)

	if result.Error != nil {
		return 0, storage.ErrUserCreateFailed
	}

	return user.Id, nil
}

func (s *Storage) FindUser(user *domain.Users) (int64, string, bool, error) {
	result := s.db.Where("login = ?", user.Login).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, "", false, storage.ErrUserNotFound
		}
		return 0, "", false, storage.ErrInternalError
	}
	return user.Id, user.Password, user.EnableTwoFa, nil
}
