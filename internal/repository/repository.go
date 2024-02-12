package repository

import (
	"context"
	"errors"

	"github.com/elkcityhazard/contact-keeper/internal/config"
	"github.com/elkcityhazard/contact-keeper/internal/models"
	"github.com/go-sql-driver/mysql"
)

type SqlUserRepository struct {
	app *config.AppConfig
}

type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

func NewSqlUserRepository(a *config.AppConfig) *SqlUserRepository {
	return &SqlUserRepository{
		app: a,
	}
}

func (s *SqlUserRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	// Implementation to fetch all users from SQL database
	ctx, cancel := context.WithTimeout(ctx, s.app.DBTimeout)

	defer cancel()

	userChan := make(chan []*models.User)
	errorChan := make(chan error)
	doneChan := make(chan bool)

	go func() {

		// Implementation to fetch all users from SQL database

		// query for all users from sql database

		query := `SELECT * FROM users`

		// execute query

		rows, err := s.app.DB.QueryContext(ctx, query, nil)

		// if ther are any errors with the query, return the error

		if err != nil {
			errorChan <- err
			return
		}

		// close rows when done

		defer rows.Close()

		// create an empty user slice

		var users []*models.User

		// loop through rows

		for rows.Next() {

			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.FirstName,
				&user.LastName,
				&user.Email,
				&user.Password.Hash,
				&user.CreatedAt,
				&user.UpdatedAt,
				&user.Version,
				&user.Scope,
			)

			// if there are any errors, return the error

			if err != nil {
				errorChan <- err
				return
			}

			// append the user to the users slice

			users = append(users, &user)
		}

		userChan <- users

		doneChan <- true

		<-doneChan

	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errorChan:
		return nil, err
	case users := <-userChan:
		return users, nil
	}
}

func (s *SqlUserRepository) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	// Implementation to fetch a single user by ID from SQL database

	return nil, nil
}

func (s *SqlUserRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	// Implementation to create a new user in SQL database

	ctx, cancel := context.WithTimeout(ctx, s.app.DBTimeout)

	defer cancel()

	userChan := make(chan *models.User, 1)
	errorChan := make(chan error, 1)

	go func() {
		defer close(userChan)
		defer close(errorChan)
		// Implementation to create a new user in SQL database

		query := `INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			created_at, 
			updated_at, 
			version, 
			scope) VALUES (?, ?, ?, ?, NOW(), NOW(), 1, ?)`

		args := []any{user.FirstName, user.LastName, user.Email, user.Password.Hash, user.Scope}

		result, err := s.app.DB.ExecContext(ctx, query, args...)

		if err != nil {
			if driverErr, ok := err.(*mysql.MySQLError); ok {
				if driverErr.Number == 1062 {

					errorChan <- errors.New("there has been an error")
					return
				}
			} else {

				errorChan <- err
			}
		}

		userID, err := result.LastInsertId()

		if err != nil {

			errorChan <- err
			return
		}

		user.ID = userID
		user.Version = 1

		userChan <- &user

	}()

	err := <-errorChan

	if err != nil {
		return nil, err
	}

	insertedUser := <-userChan
	return insertedUser, nil

}

func (s *SqlUserRepository) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	// Implementation to update an existing user in SQL database

	return nil, nil
}

func (s *SqlUserRepository) DeleteUser(ctx context.Context, id int64) error {
	// Implementation to delete a user from SQL database

	return nil
}
