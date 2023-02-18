package repository

type MessageRepository interface {
	AddMessage(id string) error
	GetMessages(id string) (int, error)
}
