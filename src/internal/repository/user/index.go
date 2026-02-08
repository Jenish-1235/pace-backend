package user

import (
	"context"
	"errors"

	userDom "pace-backend/src/internal/domain/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) userDom.UserRepository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *userDom.User) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		INSERT INTO users (id, name, email, career)
		VALUES ($1, $2, LOWER($3), $4)
	`,
		user.GetID(),
		user.GetName(),
		user.GetEmail(),
		nullify(user.GetCareer()),
	)
	if err != nil {
		return err
	}

	for _, interest := range user.GetInterests() {
		_, err = tx.Exec(ctx, `
			INSERT INTO user_interests (user_id, interest)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING
		`, user.GetID(), interest)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *repository) GetByID(ctx context.Context, id string) (*userDom.User, error) {
	row := &userRow{}

	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, career, created_at, updated_at, is_deleted
		FROM users
		WHERE id = $1 AND is_deleted = FALSE
	`, id).Scan(
		&row.ID,
		&row.Name,
		&row.Email,
		&row.Career,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.IsDeleted,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, userDom.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	interests, err := r.GetInterests(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return toDomainUser(row, interests)
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*userDom.User, error) {
	row := &userRow{}

	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, career, created_at, updated_at, is_deleted
		FROM users
		WHERE LOWER(email) = LOWER($1) AND is_deleted = FALSE
	`, email).Scan(
		&row.ID,
		&row.Name,
		&row.Email,
		&row.Career,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.IsDeleted,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, userDom.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	interests, err := r.GetInterests(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return toDomainUser(row, interests)
}

func (r *repository) Update(ctx context.Context, user *userDom.User) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		UPDATE users
		SET name = $1, email = LOWER($2), career = $3, updated_at = now()
		WHERE id = $4 AND is_deleted = FALSE
	`,
		user.GetName(),
		user.GetEmail(),
		nullify(user.GetCareer()),
		user.GetID(),
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `DELETE FROM user_interests WHERE user_id = $1`, user.GetID())
	if err != nil {
		return err
	}

	for _, interest := range user.GetInterests() {
		_, err = tx.Exec(ctx, `
			INSERT INTO user_interests (user_id, interest)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING
		`, user.GetID(), interest)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}


func (r *repository) SoftDelete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx, `
		UPDATE users
		SET is_deleted = TRUE, updated_at = now()
		WHERE id = $1 AND is_deleted = FALSE
	`, id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return userDom.ErrUserNotFound
	}

	return nil
}


func (r *repository) ExistsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE id = $1 AND is_deleted = FALSE
		)
	`, id).Scan(&exists)
	return exists, err
}

func (r *repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE LOWER(email) = LOWER($1) AND is_deleted = FALSE
		)
	`, email).Scan(&exists)
	return exists, err
}



func (r *repository) GetInterests(ctx context.Context, userID string) ([]string, error) {
	rows, err := r.db.Query(ctx,
		`SELECT interest FROM user_interests WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interests []string
	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
		interests = append(interests, i)
	}
	return interests, nil
}
