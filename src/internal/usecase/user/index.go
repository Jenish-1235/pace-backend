package user

import (
	"context"
	userDom "pace-backend/src/internal/domain/user"
)

type UseCase struct {
	repo userDom.UserRepository
}

func NewUseCase(repo userDom.UserRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) GetOrCreateUser(ctx context.Context, id, name, email, career string, interests []string) (*userDom.User, error) {
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if user != nil {
		return user, nil
	}
	user, err = userDom.NewUser(id, name, email, career, interests)
	if err != nil {
		return nil, err
	}
	
	err = uc.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}