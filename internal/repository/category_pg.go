package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/acronix0/REST-API-Go/internal/domain"
	"github.com/lib/pq"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {

	rows, err := r.db.QueryContext(ctx, "SELECT id, article, name, image FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Article, &category.Image)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) CreateOrUpdateBatch(ctx context.Context, categories []domain.Category) error {
	articles := make([]string, len(categories))
	for i, catgegory := range categories {
		articles[i] = catgegory.Article
	}

	existingCategories, err := r.GetByArticles(ctx, articles)
	if err != nil {
		return err
	}

	var newCategories []domain.Category
	var updateCategories []domain.Category

	for _, category := range categories {
		existingCategory, exist := existingCategories[category.Article]
		if exist {
			if existingCategory.Name != category.Name || existingCategory.Image != category.Image {
				category.ID = existingCategory.ID
				updateCategories = append(updateCategories, category)
			}
		} else {
			newCategories = append(newCategories, category)
		}
	}
	if len(newCategories) > 0 {
		if err := r.CreateBatch(ctx, newCategories); err != nil {
			return err
		}
	}
	if len(updateCategories) > 0 {
		if err := r.UpdateBatch(ctx, updateCategories); err != nil {
			return err
		}
	}
	return nil
}

func (r *categoryRepo) GetByArticles(ctx context.Context, articles []string) (map[string]domain.Category, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, article, name, image FROM categories WHERE article IN ($1)",
		pq.Array(articles),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make(map[string]domain.Category)
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Article, &category.Name, &category.Image)
		if err != nil {
			return nil, err
		}
		categories[category.Article] = category
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) CreateBatch(ctx context.Context, categories []domain.Category) error {
	if len(categories) == 0 {
    return errors.New("empty products array")
	}
	query := "INSERT INTO categories (article, name, image) VALUES "
	values := []interface{}{}
	placeholders := []string{}
	for i, category := range categories {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d)", i*4+1, i*4+2, i*4+3))
		values = append(values, category.Article, category.Name, category.Image)
	}
	query += strings.Join(placeholders, ", ")
	_, err := r.db.ExecContext(ctx, query, values...)
	return err
}

func (r *categoryRepo) UpdateBatch(ctx context.Context, categories []domain.Category) error {
	if len(categories) == 0 {
    return errors.New("empty products array")
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, category := range categories {
		_, err := tx.ExecContext(
			ctx,
			"UPDATE categories SET article = $1, name = $2, image = $3 WHERE id = $4",
			category.Article,
			category.Name,
			category.Image,
			category.ID,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
