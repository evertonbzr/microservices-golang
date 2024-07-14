package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/evertonbzr/microservices-golang/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewUserService(db *gorm.DB, cache *redis.Client) *UserService {
	return &UserService{
		DB:    db,
		Cache: cache,
	}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	res := s.Cache.Get(context.Background(), fmt.Sprintf("user:%d", id))

	user := &model.User{}

	if res.Err() != nil && res.Err() == redis.Nil {
		slog.Info("Cache miss", "user", id)

		if err := s.DB.First(user, id).Error; err != nil {
			return nil, err
		}

		b, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}

		if err := s.Cache.Set(context.Background(), fmt.Sprintf("user:%d", id), b, 0).Err(); err != nil {
			return nil, err
		}

		return user, nil
	}

	b, err := res.Bytes()

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers() ([]model.User, error) {
	users := []model.User{}

	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.DB.Create(user).Error
}
