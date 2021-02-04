package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
)

type mysqlUserAddressRepository struct {
	Conn *sql.DB
}

func NewMysqlUserAddressRepository(Conn *sql.DB) domain.UserAddressRepository {
	return &mysqlUserAddressRepository{Conn: Conn}
}

func (m *mysqlUserAddressRepository) fetchUserAddress(ctx context.Context, query string, args ...interface{}) (result *domain.UserAddress, err error) {
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

	result = new(domain.UserAddress)
	for rows.Next() {
		err = rows.Scan(
			&result.AddressID,
			&result.UserID,
			&result.Address,
			&result.City,
			&result.State,
			&result.Country,
			&result.PostalCode,
			&result.CreatedAt,
			&result.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
	}

	return result, nil
}

func (m *mysqlUserAddressRepository) FetchUserAddress(ctx context.Context, userID string) (res *domain.UserAddress, err error) {
	query := `SELECT
				address_id, user_id, address, city, state, country, postal_code, created_at, updated_at			  
			  FROM user_addresses
			  WHERE user_id = $1 LIMIT 1`
	if err != nil {
		return nil, domain.ErrBadParamInput
	}

	res, err = m.fetchUserAddress(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	return
}

func (m *mysqlUserAddressRepository) Store(ctx context.Context, userAddress domain.UserAddress) (err error) {
	query := `INSERT INTO user_addresses(address_id, user_id, address, city, state, country, postal_code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, userAddress.AddressID, userAddress.UserID, userAddress.Address, userAddress.City, userAddress.State, userAddress.Country, userAddress.PostalCode, userAddress.CreatedAt, userAddress.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (m *mysqlUserAddressRepository) Update(ctx context.Context, userAddress domain.UserAddress) (err error) {
	query := `
		UPDATE 
			user_addresses 
		SET 
			address=$1, city=$2, state=$3, country=$4, postal_code=$5, updated_at=$6
		WHERE user_id=$7`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, userAddress.Address, userAddress.City, userAddress.State, userAddress.Country, userAddress.PostalCode, userAddress.UpdatedAt, userAddress.UserID)
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
