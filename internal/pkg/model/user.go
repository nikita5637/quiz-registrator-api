package model

// User ...
type User struct {
	ID         int32
	Name       string
	TelegramID int64
	Email      string
	Phone      string
	State      UserState
	CreatedAt  DateTime
	UpdatedAt  DateTime
}
