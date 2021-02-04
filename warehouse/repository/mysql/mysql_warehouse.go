package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/warehouse/repository"
)

type mysqlWarehouseRepository struct {
	Conn *sql.DB
}

func NewMysqlWarehouseRepository(Conn *sql.DB) domain.WarehouseRepository {
	return &mysqlWarehouseRepository{Conn: Conn}
}

func (w *mysqlWarehouseRepository) fetchs(ctx context.Context, query string, args ...interface{}) (result []domain.Warehouse, err error) {
	rows, err := w.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.Warehouse, 0)
	for rows.Next() {
		warehouse := domain.Warehouse{}
		err := rows.Scan(
			&warehouse.ID,
			&warehouse.Address.AddressID,
			&warehouse.Name,
			&warehouse.Address.State,
			&warehouse.Address.Address,
			&warehouse.Address.City,
			&warehouse.Address.Country,
			&warehouse.Address.PostalCode,
			&warehouse.CreatedAt,
			&warehouse.UpdatedAt,
			&warehouse.Address.CreatedAt,
			&warehouse.Address.UpdatedAt,
		)
		warehouse.Address.WarehouseID = warehouse.ID
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, warehouse)
	}
	return
}

func (w *mysqlWarehouseRepository) fetch(ctx context.Context, query string, args ...interface{}) (result *domain.WarehouseIndependence, err error) {
	rows, err := w.Conn.QueryContext(ctx, query, args...)
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

	warehouse := new(domain.WarehouseIndependence)
	for rows.Next() {
		err := rows.Scan(
			&warehouse.ID,
			&warehouse.Name,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
	}
	return warehouse, nil
}

func (w *mysqlWarehouseRepository) Fetch(ctx context.Context, num int64) (res []domain.Warehouse, nextCursor string, err error) {
	query := `SELECT warehouses.id,
				COALESCE(wa.id, ''),
				warehouses.warehouse_name,
				COALESCE(wa.state, ''),
				COALESCE(wa.address, ''),
				COALESCE(wa.city, ''),
				COALESCE(wa.country, ''),
				COALESCE(wa.postal_code,''),
				warehouses.created_at,
				warehouses.updated_at,
				COALESCE(wa.created_at, now()),
				COALESCE(wa.updated_at, now())
			 FROM warehouses
				LEFT JOIN warehouse_addresses wa ON warehouses.id = wa.warehouse_id
			 LIMIT $1;`
	if err != nil {
		return nil, "", domain.ErrBadParamInput
	}
	res, err = w.fetchs(ctx, query, num)
	if err != nil {
		return nil, "", err
	}
	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)].CreatedAt)
	}
	return
}

func (w *mysqlWarehouseRepository) FetchRoom(ctx context.Context, warehouseID string) (res domain.WarehouseRoom, err error) {
	query := `SELECT
				r.id, warehouses.warehouse_name, warehouses.id, r.name, r.created_at, r.updated_at
				FROM warehouses JOIN rooms r on warehouses.id = r.warehouse_id
			WHERE warehouses.id = $1;`
	warehouseData, err := w.GetByID(ctx, warehouseID)
	if err != nil {
		logrus.Error(err)
		return
	}
	rows, err := w.Conn.QueryContext(ctx, query, warehouseID)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}()
	res.ID = warehouseData.ID
	res.Name = warehouseData.Name
	res.Address = warehouseData.Address
	for rows.Next() {
		room := domain.RoomBelongsToWarehouse{}
		errResult := rows.Scan(
			&room.ID,
			&room.WarehouseName,
			&room.WarehouseID,
			&room.Name,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if errResult != nil {
			logrus.Error(errResult)
			return
		}
		res.Rooms = append(res.Rooms, room)
	}
	return
}

func (w *mysqlWarehouseRepository) GetByID(ctx context.Context, warehouseID string) (res domain.Warehouse, err error) {
	query := `SELECT warehouses.id,
				COALESCE(wa.id, ''),
				warehouses.warehouse_name,
				COALESCE(wa.state, ''),
				COALESCE(wa.address, ''),
				COALESCE(wa.city, ''),
				COALESCE(wa.country, ''),
				COALESCE(wa.postal_code,''),
				warehouses.created_at,
				warehouses.updated_at,
				COALESCE(wa.created_at, now()),
				COALESCE(wa.updated_at, now())
			 FROM warehouses
				LEFT JOIN warehouse_addresses wa ON warehouses.id = wa.warehouse_id
			 WHERE warehouses.id = $1;`
	list, err := w.fetchs(ctx, query, warehouseID)
	if err != nil {
		return domain.Warehouse{}, err
	}
	if len(list) >= 1 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}
func (w *mysqlWarehouseRepository) Update(ctx context.Context, warehouse *domain.Warehouse) (err error) {
	query := `UPDATE warehouses SET warehouse_name = $1, updated_at = $2 WHERE id = $3`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, warehouse.Name, warehouse.UpdatedAt, warehouse.ID)
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
func (w *mysqlWarehouseRepository) Store(ctx context.Context, warehouse *domain.Warehouse) (id string, err error) {
	query := `INSERT INTO warehouses(id, warehouse_name, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id;`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, errQ := stmt.QueryContext(ctx, warehouse.ID, warehouse.Name, warehouse.CreatedAt, warehouse.UpdatedAt)
	err = errQ
	for res.Next() {
		err = res.Scan(&id)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}
	return
}
func (w *mysqlWarehouseRepository) Delete(ctx context.Context, warehouseID string) (err error) {
	query := "DELETE FROM warehouses WHERE id = $1"

	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, warehouseID)
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

func (w *mysqlWarehouseRepository) IsWarehouseExist(ctx context.Context, warehouseID string) (bool, error) {
	query := "SELECT id, warehouse_name FROM warehouses WHERE id = $1;"
	list, err := w.fetch(ctx, query, warehouseID)
	if err != nil {
		return false, err
	}
	if list.Name == "" && list.ID == "" {
		return false, nil
	}
	return true, nil
}
