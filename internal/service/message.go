package service

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yuno-obsessed/shikimori/internal/infra/config/logger"
)

type MessageService struct {
	db     *pgxpool.Pool
	logger logger.Logger
}

func NewMesageService(db *pgxpool.Pool) MessageService {
	return MessageService{
		db:     db,
		logger: logger.NewLogger(),
	}
}

func (ms MessageService) AddMessage(id string) error {

	query, args, err := squirrel.Update("messages").
		Set("message_count", squirrel.Expr("message_count + 1")).
		Where(squirrel.Eq{"user_id": id}).ToSql()
	if err != nil {
		ms.logger.Error(err.Error())
		return err
	}

	_, err = ms.db.Exec(context.Background(), query, args)
	if err != nil {
		ms.logger.Error(err.Error())
		return err
	}

	return nil
}
func (ms MessageService) GetMessages(id string) (int, error) {
	var result int

	query, args, err := squirrel.Select("message_count").
		From("messages").Where(squirrel.Eq{"user_id": id}).ToSql()
	if err != nil {
		ms.logger.Error(err.Error())
		return 0, err
	}

	row := ms.db.QueryRow(context.Background(), query, args)
	if err = row.Scan(&result); err != nil {
		ms.logger.Error(err.Error())
		return 0, err
	}

	return result, nil
}
