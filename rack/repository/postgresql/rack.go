package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/rack/repository"
)

type postgresqlRackRepository struct {
	Conn *sql.DB
}

// NewpostgresqlRackRepository will create an object that represent the article.Repository interface
func NewPostgresqlRackRepository(Conn *sql.DB) domain.RackRepository {
	return &postgresqlRackRepository{Conn}
}

func (m *postgresqlRackRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Rack, err error) {
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

	result = make([]domain.Rack, 0)
	for rows.Next() {
		t := domain.Rack{}
		err = rows.Scan(
			&t.ID,
			&t.RoomID,
			&t.Name,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *postgresqlRackRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Rack, nextCursor string, err error) {
	query := `SELECT id, room_id, name, created_at, updated_at
  						FROM racks WHERE created_at > $1 ORDER BY created_at LIMIT $2 `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}
func (m *postgresqlRackRepository) GetByID(ctx context.Context, id string) (res domain.Rack, err error) {
	query := `SELECT id, room_id, name, updated_at, created_at FROM racks WHERE id = $1`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Rack{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *postgresqlRackRepository) GetByRoomID(ctx context.Context, roomID string) ([]domain.RackDetail, error) {
	query := `SELECT id, room_id, name, updated_at, created_at FROM racks WHERE room_id = $1`

	rows, err := m.Conn.QueryContext(ctx, query, roomID)
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
	list := make([]domain.RackDetail, 0)
	for rows.Next() {
		t := domain.RackDetail{}
		err = rows.Scan(
			&t.ID,
			&t.RoomID,
			&t.Name,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		list = append(list, t)
	}
	if err != nil {
		return []domain.RackDetail{}, err
	}
	return list, nil
}

func (m *postgresqlRackRepository) Store(ctx context.Context, rack *domain.Rack) (err error) {
	query := `INSERT INTO racks(id, room_id, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, rack.ID, rack.RoomID, rack.Name, rack.UpdatedAt, rack.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (m *postgresqlRackRepository) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM racks WHERE id = $1"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
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
func (m *postgresqlRackRepository) Update(ctx context.Context, rack *domain.Rack) (err error) {
	query := `UPDATE racks set name=$1, updated_at=$2 WHERE id = $3`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, rack.Name, rack.UpdatedAt, rack.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird Behavior. Total Affected: %d", affect)
		return
	}

	return
}
