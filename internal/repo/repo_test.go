package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	db2 "github.com/ozonva/ova-entertainment-api/internal/db"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/pressly/goose"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var r Repo

var model1 = &models.Entertainment{
	ID:          uint64(1),
	UserID:      1,
	Title:       "Title 1",
	Description: "Description 1",
}

var model2 = &models.Entertainment{
	ID:          uint64(2),
	UserID:      2,
	Title:       "Title 2",
	Description: "Description 2",
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	err = resource.Expire(120)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var dbConnect *sqlx.DB
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		dbConnect, err = db2.Connect(databaseUrl)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = goose.Run("up", dbConnect.DB, "../../migration")
	if err != nil {
		log.Fatalf("Goose say krya: %s", err)
	}

	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		r = NewRepo(dbConnect)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		dbConnect.Close()
	}()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreate(t *testing.T) {
	err := r.AddEntertainments([]models.Entertainment{*model1})
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	model2.Title = "Title New"
	_, err := r.UpdateEntertainment(*model2)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	err := r.RemoveEntertainment(model1.ID)
	assert.NoError(t, err)
}
