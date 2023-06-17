package model

// Role ...
type Role uint8

const (
	// RoleInvalid ...
	RoleInvalid Role = 0
	// RoleAdmin ...
	RoleAdmin Role = 1
	// RoleManagement ...
	RoleManagement Role = 2
	// RoleUser ...
	RoleUser Role = 3

	invalid    = "invalid"
	admin      = "admin"
	management = "management"
	user       = "user"
)

// String ...
func (r Role) String() string {
	switch r {
	case RoleInvalid:
		return invalid
	case RoleAdmin:
		return admin
	case RoleManagement:
		return management
	case RoleUser:
		return user
	}

	return invalid
}

// UserRole ...
type UserRole struct {
	ID     int32
	UserID int32
	Role   Role
}
