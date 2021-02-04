package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/user/repository"
)

type postgresqlUserRepository struct {
	Conn *sql.DB
}

func NewPostgresqlUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresqlUserRepository{Conn: Conn}
}

func (m *postgresqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Address.AddressID,
			&t.Email,
			&t.Password,
			&t.FirstName,
			&t.LastName,
			&t.Gender,
			&t.DateOfBirth,
			&t.Address.Address,
			&t.Address.City,
			&t.Address.State,
			&t.Address.Country,
			&t.Address.PostalCode,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		t.Address.CreatedAt = t.CreatedAt
		t.Address.UpdatedAt = t.UpdatedAt
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresqlUserRepository) Fetch(ctx context.Context, num int64) (res []domain.User, nextCursor string, err error) {
	query := `SELECT
    				users.id, ua.address_id, email, password, first_name, last_name, gender, date_of_birth, address, city, state, country, postal_code, users.created_at, users.updated_at
				FROM users
					JOIN user_addresses ua on users.id = ua.user_id LIMIT $1;`
	if err != nil {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, num)
	if err != nil {
		return nil, "", err
	}
	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)].CreatedAt)
	}
	return
}

func (m *postgresqlUserRepository) GetByID(ctx context.Context, userID string) (res domain.User, err error) {
	query := `SELECT
	users.id, ua.address_id, email, password, first_name, last_name, gender, date_of_birth, address, city, state, country, postal_code, users.created_at, users.updated_at
  			  FROM users
				JOIN user_addresses ua on users.id = ua.user_id 
			  WHERE users.id = $1`

	list, err := m.fetch(ctx, query, userID)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *postgresqlUserRepository) Store(ctx context.Context, user *domain.User) (err error) {
	insertUserQuery := `INSERT INTO users(id, email, password, first_name, last_name, gender, date_of_birth, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := m.Conn.PrepareContext(ctx, insertUserQuery)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, user.ID, user.Email, user.Password, user.FirstName, user.LastName, user.Gender, user.DateOfBirth, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (m *postgresqlUserRepository) Delete(ctx context.Context, userID string) (err error) {
	query := "DELETE FROM users WHERE id = $1"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, userID)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *postgresqlUserRepository) Update(ctx context.Context, user *domain.User) (err error) {
	query := `UPDATE users SET email=$1, first_name=$2, last_name=$3, gender=$4, date_of_birth=$5, updated_at=$6 WHERE id = $7`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, user.Email, user.FirstName, user.LastName, user.Gender, user.DateOfBirth, user.UpdatedAt, user.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
