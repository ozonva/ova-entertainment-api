package repo

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-entertainment-api/internal/models"
)

type Repo interface {
	AddEntertainments(models []models.Entertainment) error
	ListEntertainments(limit uint32, offset uint32) ([]models.Entertainment, error)
	DescribeEntertainment(model models.Entertainment) (*models.Entertainment, error)
	RemoveEntertainment(ID uint64) error
	FindEntertainment(ID uint64, title string) (*models.Entertainment, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{db: db}
}

func (r *repo) AddEntertainments(models []models.Entertainment) error {

	query := "INSERT INTO entertainments (user_id, title, description) VALUES (:user_id, :title, :description)"
	_, err := r.db.NamedExec(query, models)

	return err
}

func (r *repo) ListEntertainments(limit uint32, offset uint32) ([]models.Entertainment, error) {

	result := make([]models.Entertainment, 0, limit)

	query := "SELECT * FROM entertainments ORDER BY id DESC LIMIT $1 OFFSET $2"
	rows, err := r.db.Queryx(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var entertainment models.Entertainment
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
	query := "UPDATE entertainments SET title = :title, description = :description WHERE id = :id"
	result, err := r.db.NamedExec(query, model)
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
	query := "DELETE FROM entertainments WHERE id = $1"
	_, err := r.db.Exec(query, ID)

	return err
}

func (r *repo) FindEntertainment(userID uint64, title string) (*models.Entertainment, error) {
	query := "SELECT * FROM entertainments WHERE user_id = $1 and title = $2"
	row := r.db.QueryRowx(query, userID, title)
	entertainment := models.Entertainment{}

	err := row.StructScan(&entertainment)
	if err == sql.ErrNoRows {
		return nil, errors.New("Not found")
	} else if err != nil {
		return nil, err
	}

	return &entertainment, nil
}
