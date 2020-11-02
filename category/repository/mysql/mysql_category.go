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
  						FROM categories WHERE created_at > $1 ORDER BY created_at LIMIT $2 `

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
func (m *mysqlCategoryRepository) GetByID(ctx context.Context, id string) (res domain.Category, err error) {
	query := `SELECT id, category, updated_at, created_at FROM categories WHERE ID = $1`

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
	query := `INSERT INTO categories(id, category, updated_at, created_at) VALUES($1, $2, $3, $4)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, category.ID, category.Category, category.UpdatedAt, category.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (m *mysqlCategoryRepository) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM categories WHERE id = $1"

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
func (m *mysqlCategoryRepository) Update(ctx context.Context, category *domain.Category) (err error) {
	query := `UPDATE categories set category=$1, updated_at=$2 WHERE ID = $3`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, category.Category, category.UpdatedAt, category.ID)
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
