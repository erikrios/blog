package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/erikrios/blog/config"
	"github.com/erikrios/blog/internal/entity"
	"github.com/erikrios/blog/internal/repository"
	"github.com/erikrios/blog/internal/util"
)

func createConnection(t *testing.T) *sql.DB {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		t.FailNow()
	}

	if err := db.Ping(); err != nil {
		t.FailNow()
	}

	t.Cleanup(
		func() {
			res, err := db.Exec("DELETE FROM users;")
			if err != nil {
				t.Fatal(err)
			}

			count, err := res.RowsAffected()
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Cleaning up the database, %d records from users table deleted.\n", count)
			_ = db.Close()
		},
	)

	return db
}

func TestInsert(t *testing.T) {
	db := createConnection(t)
	var repo repository.UserRepository = repository.NewUserRepositoryImpl(db)

	t.Run("it should return error, when username is duplicated", func(t *testing.T) {
		user1 := insertUser(t, repo)

		user2 := entity.User{
			Username: user1.Username,
			Name:     util.GenerateString(4) + " " + util.GenerateString(5),
			Password: util.GenerateString(8),
		}

		_, err := repo.Insert(context.Background(), user2)
		if err == nil {
			t.FailNow()
		}

		if !errors.Is(err, repository.ErrAlreadyExists) {
			t.FailNow()
		}
	})

	t.Run("it should successfully inserted, when the user is valid", func(t *testing.T) {
		insertUser(t, repo)
	})
}

func insertUser(t testing.TB, repo repository.UserRepository) entity.User {
	user := entity.User{
		Username: util.GenerateString(8),
		Name:     util.GenerateString(4) + " " + util.GenerateString(5),
		Password: util.GenerateString(8),
	}

	var err error
	user.ID, err = repo.Insert(context.Background(), user)
	if err != nil {
		t.Error(err)
	}

	return user
}
