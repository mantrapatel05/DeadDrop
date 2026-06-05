package server

import (
	"context"
	"deadrop/model"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

var errNotFound = fmt.Errorf("not found")

func (s *Server) CreateDrop(ctx context.Context, drop *model.Drop) error {
	query := `
	INSERT INTO drops (id, ciphertext, reveal_at, knock_target, knock_count, burned, created_at)
	VALUES ($1, $2, $3, $4, COALESCE($5, 0), COALESCE($6, false), NOW())
	`

	_, err := s.db.Pool.Exec(ctx, query,
		drop.ID,
		drop.Ciphertext,
		drop.RevealAt,
		drop.KnockTarget,
		drop.KnockCount,
		drop.Burned,
	)
	if err != nil {
		return fmt.Errorf("create drop: %w", err)
	}
	return nil
}

func (s *Server) GetDrop(ctx context.Context, id string) (model.Drop, error) {
	query := `
	SELECT id,ciphertext,reveal_at,knock_target,knock_count,burned,created_at
	FROM drops
	WHERE id = $1`

	var drop model.Drop
	err := s.db.Pool.QueryRow(ctx, query, id).Scan(
		&drop.ID,
		&drop.Ciphertext,
		&drop.RevealAt,
		&drop.KnockTarget,
		&drop.KnockCount,
		&drop.Burned,
		&drop.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Drop{}, errNotFound
		}
		return model.Drop{}, fmt.Errorf("get drop by id: %w", err)
	}

	return drop, nil
}

func (s *Server) IncrementDrop(ctx context.Context, id string) error {
	query := `
	UPDATE drops
	SET knock_count = knock_count + 1
	WHERE id = $1
	`
	result, err := s.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("increment knock: %w", err)
	}
	if result.RowsAffected() == 0 {
		return errNotFound
	}
	return nil
}

func (s *Server) MarkBurned(ctx context.Context, id string) error {
	query := `
	UPDATE drops
	SET burned = true
	WHERE id = $1
	`
	result, err := s.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("mark burned: %w", err)
	}
	if result.RowsAffected() == 0 {
		return errNotFound
	}
	return nil
}

func (s *Server) UpdateReveal(ctx context.Context, id string, revealAt *time.Time, knockTarget *int) error {
	query := `
	UPDATE drops
	SET reveal_at = $1, knock_target = $2
	WHERE id = $3
	`

	result, err := s.db.Pool.Exec(ctx, query, revealAt, knockTarget, id)
	if err != nil {
		return fmt.Errorf("update reveal: %w", err)
	}
	if result.RowsAffected() == 0 {
		return errNotFound
	}
	return nil
}
