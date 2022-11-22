package repository

type ChatRepository interface {
	GetUser(userID string) (string, error)
}
