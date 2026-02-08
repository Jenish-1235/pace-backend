package user

import (
	userDom "pace-backend/src/internal/domain/user"
)

func toDomainUser(row *userRow, interests []string) (*userDom.User, error) {
	return userDom.NewUser(
		row.ID,
		row.Name,
		row.Email,
		deref(row.Career),
		interests,
	)
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func nullify(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}