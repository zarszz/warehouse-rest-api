package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
)

type postgresqlWarehouseAddressRepository struct {
	Conn *sql.DB
}

func NewPostgresqlWarehouseAddressRepository(Conn *sql.DB) domain.WarehouseAddressRepository {
	return &postgresqlWarehouseAddressRepository{Conn: Conn}
}

func (m *postgresqlWarehouseAddressRepository) fetchWarehouseAddress(ctx context.Context, query string, args ...interface{}) (result *domain.WarehouseAddress, err error) {
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

	result = new(domain.WarehouseAddress)
	for rows.Next() {
		err = rows.Scan(
			&result.AddressID,
			&result.WarehouseID,
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

func (m *postgresqlWarehouseAddressRepository) FetchWarehouseAddress(ctx context.Context, userID string) (res *domain.WarehouseAddress, err error) {
	query := `SELECT 
				id, warehouse_id, address, province, city, country, state, postal_code, created_at, updated_at
			  FROM warehouse_addresses WHERE warehouse_id = $1 LIMIT 1`
	if err != nil {
		return nil, domain.ErrBadParamInput
	}

	res, err = m.fetchWarehouseAddress(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	return
}

func (m *postgresqlWarehouseAddressRepository) Store(ctx context.Context, warehouseAddress domain.WarehouseAddress) (err error) {
	query := `INSERT INTO warehouse_addresses(id, warehouse_id, address, city, state, country, postal_code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, warehouseAddress.AddressID, warehouseAddress.WarehouseID, warehouseAddress.Address, warehouseAddress.City, warehouseAddress.State, warehouseAddress.Country, warehouseAddress.PostalCode, warehouseAddress.CreatedAt, warehouseAddress.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (m *postgresqlWarehouseAddressRepository) Update(ctx context.Context, warehouseAddress domain.WarehouseAddress) (err error) {
	query := `
		UPDATE 
			warehouse_addresses 
		SET 
			address=$1, city=$2, state=$3, country=$4, postal_code=$5, updated_at=$6
		WHERE warehouse_id=$7`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, warehouseAddress.Address, warehouseAddress.City, warehouseAddress.State, warehouseAddress.Country, warehouseAddress.PostalCode, warehouseAddress.UpdatedAt, warehouseAddress.WarehouseID)
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
