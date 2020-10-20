package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/zarszz/warehouse-rest-api/category/repository"
	"github.com/zarszz/warehouse-rest-api/domain"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

// NewMysqlCategoryRepository will create an object that represent the article.Repository interface
func NewMysqlCategoryRepository(Conn *sql.DB) domain.CategoryRepository {
	return &mysqlCategoryRepository{Conn}
}

func (m *mysqlCategoryRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Category, err error) {
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

	result = make([]domain.Category, 0)
	for rows.Next() {
		t := domain.Category{}
		err = rows.Scan(
			&t.ID,
			&t.Category,
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

func (m *mysqlCategoryRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Category, nextCursor string, err error) {
	query := `SELECT id,category,created_at, updated_at
  						FROM category WHERE created_at > ? ORDER BY created_at LIMIT ? `

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
func (m *mysqlCategoryRepository) GetByID(ctx context.Context, id int64) (res domain.Category, err error) {
	query := `SELECT id, category, updated_at, created_at FROM category WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Category{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlCategoryRepository) Store(ctx context.Context, category *domain.Category) (err error) {
	query := `INSERT INTO category SET category=? , id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, category.Category, category.ID, category.UpdatedAt, category.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (m *mysqlCategoryRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM category WHERE id = ?"

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
func (m *mysqlCategoryRepository) Update(ctx context.Context, ar *domain.Category) (err error) {
	query := `UPDATE category set category=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Category, ar.UpdatedAt, ar.ID)
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
