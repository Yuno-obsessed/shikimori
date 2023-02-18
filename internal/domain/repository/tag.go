package repository

import (
	"github.com/yuno-obsessed/shikimori/internal/domain/entity"
)

type TagRepository interface {
	AddTag(tag entity.Tag) error
	DeleteTag(id int) error
}
