package repository

import "github.com/yuno-obsessed/shikimori/internal/domain/entity"

type UserRepository interface {
	GetUser(id string) (entity.User, error)
	ChangeUsername(id string, username string) error
	UpdateRole(id string) error
	LevelUp(id string) error
}
