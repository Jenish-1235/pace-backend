package user

import "time"

type userRow struct {
	ID		string
	Name	string
	Email	string
	Career	*string
	Interests []string
	
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}

type interests struct {
	Interest string
}