package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/room/repository"
)

type postgresqlRoomRepository struct {
	Conn *sql.DB
}

func NewPostgresqlRoomRepositoryWarehouseRepository(Conn *sql.DB) domain.RoomRepository {
	return &postgresqlRoomRepository{Conn: Conn}
}

func (w *postgresqlRoomRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Room, err error) {
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

	result = make([]domain.Room, 0)
	for rows.Next() {
		room := domain.Room{}
		err := rows.Scan(
			&room.ID,
			&room.WarehouseID,
			&room.Name,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, room)
	}
	return
}

func (w *postgresqlRoomRepository) Fetch(ctx context.Context, num int64) (res []domain.Room, nextCursor string, err error) {
	query := `SELECT id, warehouse_id, name, created_at, updated_at FROM rooms LIMIT $1`
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
func (w *postgresqlRoomRepository) GetByID(ctx context.Context, roomID string) (res domain.Room, err error) {
	query := `SELECT id, warehouse_id, name, created_at, updated_at FROM rooms WHERE id = $1`
	list, err := w.fetch(ctx, query, roomID)
	if err != nil {
		return domain.Room{}, err
	}
	if len(list) >= 1 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (w *postgresqlRoomRepository) GetByWarehouseID(ctx context.Context, warehouseID string) (res []domain.RoomDetail, err error) {
	query := `SELECT id, warehouse_id, name, created_at, updated_at FROM rooms WHERE warehouse_id = $1`
	rows, err := w.Conn.QueryContext(ctx, query, warehouseID)
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

	for rows.Next() {
		room := domain.RoomDetail{}
		err := rows.Scan(
			&room.ID,
			&room.WarehouseID,
			&room.Name,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, room)
	}
	if len(res) == 0 {
		return []domain.RoomDetail{}, nil
	}
	return
}

func (w *postgresqlRoomRepository) Update(ctx context.Context, room *domain.Room) (err error) {
	query := `UPDATE rooms SET name = $1, updated_at = $2 WHERE id = $3`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, room.Name, room.UpdatedAt, room.ID)
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
func (w *postgresqlRoomRepository) Store(ctx context.Context, room *domain.Room) (err error) {
	query := `INSERT INTO rooms(id, warehouse_id, name, created_at, updated_at) VALUES($1, $2, $3, $4, $5);`
	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx, room.ID, room.WarehouseID, room.Name, room.CreatedAt, room.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func (w *postgresqlRoomRepository) Delete(ctx context.Context, roomID string) (err error) {
	query := "DELETE FROM rooms WHERE id = $1"

	stmt, err := w.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, roomID)
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

func (w *postgresqlRoomRepository) IsRoomExist(ctx context.Context, roomID string) (bool, error) {
	query := "SELECT * FROM rooms WHERE id = $1"
	list, err := w.fetch(ctx, query, roomID)
	if err != nil {
		return false, err
	}
	if len(list) <= 0 {
		return false, nil
	}
	return true, nil
}
