package repo

import (
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-entertainment-api/internal/models"
)

type Repo interface {
	AddEntertainments(models []models.Entertainment) error
	ListEntertainments(limit uint32, offset uint32) ([]models.Entertainment, error)
	DescribeEntertainment(model models.Entertainment) (*models.Entertainment, error)
	RemoveEntertainment(ID uint64) error
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) AddEntertainments(models []models.Entertainment) error {

	builder := squirrel.
		Insert("entertainments").
		Columns("user_id", "title", "description")

	for _, model := range models {
		builder = builder.Values(model.UserID, model.Title, model.Description)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)

	return err
}

func (r *repo) ListEntertainments(limit uint32, offset uint32) ([]models.Entertainment, error) {

	var entertainment models.Entertainment
	result := make([]models.Entertainment, 0, limit)

	query, args, err := squirrel.
		Select("*").
		From("entertainments").
		OrderBy("id DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&entertainment); err != nil {
			return nil, err
		}
		result = append(result, entertainment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo) DescribeEntertainment(model models.Entertainment) (*models.Entertainment, error) {
	query, args, err := squirrel.
		Update("entertainments").
		Set("title", model.Title).
		Set("description", model.Description).
		Where(squirrel.Eq{"id": model.ID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return nil, errors.New("Nothing updated")
	}

	return &model, nil
}

func (r *repo) RemoveEntertainment(ID uint64) error {
	query, args, err := squirrel.
		Delete("entertainments").
		Where(squirrel.Eq{"id": ID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)

	return err
}
