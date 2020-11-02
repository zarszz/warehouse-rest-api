package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/category/repository"
	"github.com/zarszz/warehouse-rest-api/domain"
)

type mysqlWarehouseRepository struct {
	Conn *sql.DB
}

func NewMysqlWarehouseRepository(Conn *sql.DB) domain.WarehouseRepository {
	return &mysqlWarehouseRepository{Conn: Conn}
}

func (w *mysqlWarehouseRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Warehouse, err error) {
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
			&warehouse.Name,
			&warehouse.CreatedAt,
			&warehouse.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, warehouse)
	}
	return
}

func (w *mysqlWarehouseRepository) Fetch(ctx context.Context, num int64) (res []domain.Warehouse, nextCursor string, err error) {
	query := `SELECT id, warehouse_name, created_at, updated_at FROM warehouses LIMIT $1`
	if err != nil {
		return nil, "", domain.ErrBadParamInput
	}
	res, err = w.fetch(ctx, query, num)
	if err != nil {
		return nil, "", err
	}
	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)].CreatedAt)
	}
	return
}
func (w *mysqlWarehouseRepository) GetByID(ctx context.Context, warehouseID string) (res domain.Warehouse, err error) {
	query := `SELECT id, warehouse_name, created_at, updated_at FROM warehouses WHERE id = $1;`
	list, err := w.fetch(ctx, query, warehouseID)
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
func (w *mysqlWarehouseRepository) Store(ctx context.Context, warehouse *domain.Warehouse) (err error) {
	query := `INSERT INTO warehouses(id, warehouse_name, created_at, updated_at) VALUES($1, $2, $3, $4);`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, warehouse.ID, warehouse.Name, warehouse.CreatedAt, warehouse.UpdatedAt)
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
