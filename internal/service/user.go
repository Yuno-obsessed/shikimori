package service

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yuno-obsessed/shikimori/internal/domain/entity"
	"github.com/yuno-obsessed/shikimori/internal/infra/config/logger"
)

type UserService struct {
	db  *pgxpool.Pool
	log logger.Logger
}

func NewUserService(db *pgxpool.Pool) UserService {
	return UserService{
		db:  db,
		log: logger.NewLogger(),
	}
}

func (us UserService) GetUser(id string) (entity.User, error) {
	var user entity.User

	query, args, err := squirrel.Select("user_id",
		"username", "user_level", "role_id").From("users").
		Where(squirrel.Eq{"user_id": id}).ToSql()
	if err != nil {
		us.log.Error(err.Error())
		return entity.User{}, err
	}

	row := us.db.QueryRow(context.Background(), query, args...)
	if err := row.Scan(&user.Id, &user.Username,
		&user.Level, &user.RoleId); err != nil {
		us.log.Error(err.Error())
		return entity.User{}, err
	}

	return user, nil
}

func (us UserService) ChangeUsername(id string, username string) error {

	query, args, err := squirrel.Update("users").
		Set("username", username).Where(squirrel.Eq{"role_id": id}).ToSql()
	if err != nil {
		us.log.Error(err.Error())
		return err
	}

	_, err = us.db.Exec(context.Background(), query, args)
	if err != nil {
		us.log.Error(err.Error())
		return err
	}
	return nil
}
func (us UserService) UpdateRole(id string) error {

	query, args, err := squirrel.Update("users").
		Set("role_id", squirrel.Expr("role_id+1")).
		Where(squirrel.Eq{"user_id": id}).ToSql()
	if err != nil {
		us.log.Error(err.Error())
		return err
	}

	_, err = us.db.Exec(context.Background(), query, args)
	if err != nil {
		us.log.Error(err.Error())
		return err
	}

	return nil
}
func (us UserService) LevelUp(id string) error {

	query, args, err := squirrel.Update("users").
		Set("level", squirrel.Expr("level+1")).
		Where(squirrel.Eq{"user_id": id}).ToSql()
	if err != nil {
		us.log.Error(err.Error())
		return err
	}

	_, err = us.db.Exec(context.Background(), query, args)
	if err != nil {
		us.log.Error(err.Error())
		return err
	}

	return nil
}
