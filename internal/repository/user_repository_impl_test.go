package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"sync"
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
			t.Fatalf("want %v, got %v", repository.ErrAlreadyExists, err)
		}

		if !errors.Is(err, repository.ErrAlreadyExists) {
			t.Fatalf("want %v, got %v", repository.ErrAlreadyExists, err)
		}
	})

	t.Run("it should successfully inserted, when the user is valid", func(t *testing.T) {
		insertUser(t, repo)
	})
}

func TestFindAll(t *testing.T) {
	db := createConnection(t)
	var repo repository.UserRepository = repository.NewUserRepositoryImpl(db)

	t.Run("it should successfully find all users, when no error occurs", func(t *testing.T) {
		const n = 100
		expectedUsers := make(map[int64]entity.User, n)

		wg := sync.WaitGroup{}
		userCh := make(chan entity.User)

		wg.Add(1)
		go func(wg *sync.WaitGroup, n int, ch chan entity.User) {
			for len(expectedUsers) != n {
			}
			close(ch)
			wg.Done()
		}(&wg, n, userCh)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, ch chan<- entity.User) {
				ch <- insertUser(t, repo)
				wg.Done()
			}(&wg, userCh)
		}

	loop:
		for {
			select {
			case user, ok := <-userCh:
				if !ok {
					t.Log("channel closed")
					break loop
				}
				t.Log("successfully insert user:", user)
				expectedUsers[user.ID] = user
			}
		}

		wg.Wait()
		gotUsers, err := repo.FindAll(context.Background())
		if err != nil {
			t.Fatalf("want %v, got %v", nil, err)
		}

		if len(gotUsers) != n {
			t.Fatalf("want %v, got %v", n, len(gotUsers))
		}

		for i := 0; i < n; i++ {
			gotUser := gotUsers[i]
			expectedUser, ok := expectedUsers[gotUser.ID]
			if !ok {
				t.Fatalf("want %v, got %v", true, ok)
			}

			if expectedUser.ID != gotUser.ID {
				t.Fatalf("want %v, got %v", expectedUser.ID, gotUser.ID)
			}

			if expectedUser.Username != gotUser.Username {
				t.Fatalf("want %v, got %v", expectedUser.Username, gotUser.Username)
			}

			if expectedUser.Name != gotUser.Name {
				t.Fatalf("want %v, got %v", expectedUser.Name, gotUser.Name)
			}

			if expectedUser.Password != gotUser.Password {
				t.Fatalf("want %v, got %v", expectedUser.Password, gotUser.Password)
			}
		}
	})
}

func TestFindByID(t *testing.T) {
	db := createConnection(t)
	var repo repository.UserRepository = repository.NewUserRepositoryImpl(db)

	t.Run("it should return not found error, when given ID is not exists", func(t *testing.T) {
		_, err := repo.FindByID(context.Background(), int64(util.GenerateInt(100)))
		if err == nil {
			t.Fatalf("want %v, got %v", repository.ErrNotFound, err)
		}

		if !errors.Is(err, repository.ErrNotFound) {
			t.Fatalf("want %v, got %v", repository.ErrNotFound, err)
		}
	})

	t.Run("it should successfully find valid user, when given ID is exists", func(t *testing.T) {
		expectedUser := insertUser(t, repo)
		gotUser, err := repo.FindByID(context.Background(), expectedUser.ID)
		if err != nil {
			t.Fatalf("want %v, got %v", nil, err)
		}

		if expectedUser.ID != gotUser.ID {
			t.Fatalf("want %v, got %v", expectedUser.ID, gotUser.ID)
		}

		if expectedUser.Username != gotUser.Username {
			t.Fatalf("want %v, got %v", expectedUser.Username, gotUser.Username)
		}

		if expectedUser.Name != gotUser.Name {
			t.Fatalf("want %v, got %v", expectedUser.Name, gotUser.Name)
		}

		if expectedUser.Password != gotUser.Password {
			t.Fatalf("want %v, got %v", expectedUser.Password, gotUser.Password)
		}
	})
}

func TestUpdate(t *testing.T) {
	db := createConnection(t)
	var repo repository.UserRepository = repository.NewUserRepositoryImpl(db)

	t.Run("it should return error, when id is not exists", func(t *testing.T) {
		id := int64(util.GenerateInt(10))
		var user entity.User

		err := repo.Update(context.Background(), id, user)
		if err == nil {
			t.Fatalf("want %v, got %v", repository.ErrUnknown, err)
		}

		if !errors.Is(err, repository.ErrUnknown) {
			t.Fatalf("want %v, got %v", repository.ErrUnknown, err)
		}
	})

	t.Run("it should successfully updated, when id is exists", func(t *testing.T) {
		expectedUser := insertUser(t, repo)
		expectedUser.Username = util.GenerateString(8)
		expectedUser.Name = util.GenerateString(4) + " " + util.GenerateString(5)
		expectedUser.Password = util.GenerateString(8)

		err := repo.Update(context.Background(), expectedUser.ID, expectedUser)

		if err != nil {
			t.Fatalf("want %v, got %v", nil, err)
		}
	})
}

func TestDelete(t *testing.T) {
	db := createConnection(t)
	var repo repository.UserRepository = repository.NewUserRepositoryImpl(db)

	t.Run("it should return error, when id is not exists", func(t *testing.T) {
		id := int64(util.GenerateInt(10))

		err := repo.Delete(context.Background(), id)
		if err == nil {
			t.Fatalf("want %v, got %v", repository.ErrUnknown, err)
		}

		if !errors.Is(err, repository.ErrUnknown) {
			t.Fatalf("want %v, got %v", repository.ErrUnknown, err)
		}
	})

	t.Run("it should successfully deleted, when id is exists", func(t *testing.T) {
		expectedUser := insertUser(t, repo)

		err := repo.Delete(context.Background(), expectedUser.ID)

		if err != nil {
			t.Fatalf("want %v, got %v", nil, err)
		}

		_, err = repo.FindByID(context.Background(), expectedUser.ID)
		if err == nil {
			t.Fatalf("want %v, got %v", repository.ErrNotFound, err)
		}

		if !errors.Is(err, repository.ErrNotFound) {
			t.Fatalf("want %v, got %v", repository.ErrNotFound, err)
		}
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
		t.Fatalf("want %v, got %v", nil, err)
	}

	return user
}
