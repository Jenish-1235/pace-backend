package user

import (
	"time"
)


type User struct {
	id 	 string
	name string
	email string
	career string
	interests []string

	createdAt time.Time
	updatedAt time.Time

	isDeleted bool
}

func NewUser(id, name, email, career string, interests []string) (*User, error) {

	if interests == nil {
		interests = []string{}
	}

	if email == "" {
		return nil, ErrInvalidUserEmail
	}

	if id == "" {
		return nil, ErrInvalidUserId
	}

	now := time.Now().UTC()
	return &User{
		id:        id,
		name:      name,
		email:     email,
		career:    career,
		interests: interests,
		createdAt: now,
		updatedAt: now,
		isDeleted: false,
	}, nil
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}	
func (u *User) GetCareer() string {
	return u.career
}
func (u *User) GetInterests() []string {
	return u.interests
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) IsDeleted() bool {
	return u.isDeleted
}

func (u *User) MarkAsDeleted() error {
	if u.isDeleted {
		return ErrDeletedUserAccess
	}
	u.isDeleted = true
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *User) Restore() {
	u.isDeleted = false
	u.updatedAt = time.Now().UTC()
}

func (u *User) UpdateName(name string) error {
	if u.isDeleted {
		return ErrDeletedUserAccess
	}
	u.name = name
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateEmail(email string) error {
	if u.isDeleted {
		return ErrDeletedUserAccess
	}
	if email == "" {
		return ErrInvalidUserEmail
	}
	u.email = email
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateCareer(career string) error {
	if u.isDeleted {
		return ErrDeletedUserAccess
	}
	u.career = career
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *User) UpdateInterests(interests []string) error {
	if u.isDeleted {
		return ErrDeletedUserAccess
	}
	if interests == nil {
		interests = []string{}
	}
	u.interests = interests
	u.updatedAt = time.Now().UTC()
	return nil
}