package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/item/repository"
)

type mysqlItemRepository struct {
	Conn *sql.DB
}

func NewMysqlItemRepository(Conn *sql.DB) domain.ItemRepository {
	return &mysqlItemRepository{Conn: Conn}
}

func (w *mysqlItemRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Item, err error) {
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

	result = make([]domain.Item, 0)
	for rows.Next() {
		item := domain.Item{}
		err := rows.Scan(
			&item.ID,
			&item.ItemName,
			&item.Description,
			&item.OwnerID,
			&item.RackID,
			&item.CategoryID,
			&item.RoomID,
			&item.WarehouseID,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, item)
	}
	return
}

func (w *mysqlItemRepository) Fetch(ctx context.Context, num int64) (res []domain.Item, nextCursor string, err error) {
	query := `SELECT id, item_name, description, owner_id, rack_id, category_id, room_id, warehouse_id, created_at, updated_at FROM items LIMIT $1`
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
func (w *mysqlItemRepository) GetByID(ctx context.Context, itemID string) (res domain.Item, err error) {
	query := `SELECT id, item_name, description, owner_id, rack_id, category_id, room_id, warehouse_id, created_at, updated_at FROM items WHERE id = $1;`
	list, err := w.fetch(ctx, query, itemID)
	if err != nil {
		return domain.Item{}, err
	}
	if len(list) >= 1 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (w *mysqlItemRepository) GetByRackID(ctx context.Context, rackID string) ([]domain.Item, error) {
	query := `SELECT id, item_name, description, owner_id, rack_id, category_id, room_id, warehouse_id, created_at, updated_at FROM items WHERE rack_id = $1;`
	res, err := w.fetch(ctx, query, rackID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return res, nil
}

func (w *mysqlItemRepository) Update(ctx context.Context, item *domain.Item) (err error) {
	query := `UPDATE items SET item_name = $1, description = $2, updated_at = $3 WHERE id = $4`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, item.ItemName, item.Description, item.UpdatedAt, item.ID)
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
func (w *mysqlItemRepository) Store(ctx context.Context, item *domain.Item) (err error) {
	query := `INSERT INTO 
    			items(id, item_name, description, owner_id, rack_id, category_id, room_id, warehouse_id, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, item.ID, item.ItemName, item.Description, item.OwnerID, item.RackID, item.CategoryID, item.RoomID, item.WarehouseID, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func (w *mysqlItemRepository) Delete(ctx context.Context, itemID string) (err error) {
	query := "DELETE FROM items WHERE id = $1"

	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, itemID)
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
