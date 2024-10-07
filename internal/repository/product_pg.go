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

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepo {
	return &productRepo{db: db}
}

func (r *productRepo) GetProducts(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, article, name, image, price FROM products",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.ID,
			&product.Article,
			&product.Name,
			&product.Image,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepo) GetByCredentials(ctx context.Context, query domain.GetProductsQuery) ([]domain.Product, error) {
	var sqlQuery strings.Builder
	var params []interface{}
	sqlQuery.WriteString("SELECT id, article, name, image, price FROM products WHERE 1=1")

	if query.Search != "" {
		sqlQuery.WriteString(" AND name ILIKE ?")
		params = append(params, "%"+query.Search+"%")
	}
	if query.MinPrice > 0 {
		sqlQuery.WriteString(" AND price >= ?")
		params = append(params, query.MinPrice)
	}
	if query.MaxPrice > 0 {
		sqlQuery.WriteString(" AND price <= ?")
		params = append(params, query.MaxPrice)
	}
	if query.InStock {
		sqlQuery.WriteString(" AND stock_count > 0")
	}
	switch query.SortedType {
	case domain.Name:
		sqlQuery.WriteString(" ORDER BY name ASC")
	case domain.PriceUp:
		sqlQuery.WriteString(" ORDER BY price ASC")
	case domain.PriceDown:
		sqlQuery.WriteString(" ORDER BY price DESC")
	case domain.Count:
		sqlQuery.WriteString(" ORDER BY stock_count DESC")
	}

	if query.Limit > 0 {
		sqlQuery.WriteString(" LIMIT ?")
		params = append(params, query.Limit)
	}
	if query.Skip > 0 {
		sqlQuery.WriteString(" OFFSET ?")
		params = append(params, query.Skip)
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery.String(), params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.ID,
			&product.Article,
			&product.Name,
			&product.Image,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

/* func (r *productRepo) GetByID(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, article, name, price, image, quantity, category_id FROM products WHERE id = $1",
		id,
	).Scan(
		&product.ID,
		&product.Article,
		&product.Name,
		&product.Price,
		&product.Image,
		&product.Quantity,
		&product.CategoryID,
	)
	if err == sql.ErrNoRows {
		return domain.Product{}, domain.ErrProductNotFound
	}
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
 */
func (r *productRepo) CreateOrUpdateBatch(ctx context.Context, products []domain.Product) error {
	articles := make([]string, len(products))
	for i, product := range products {
		articles[i] = product.Article
	}

	existingProducts, err := r.GetByArticles(ctx, articles)
	if err != nil {
		return err
	}

	var newProducts []domain.Product
	var updatedProducts []domain.Product

	for _, product := range products {
		existingProduct, exists := existingProducts[product.Article]

		if exists {
			if existingProduct.Name != product.Name || existingProduct.Price != product.Price || existingProduct.CategoryID != product.CategoryID {
				product.ID = existingProduct.ID
				updatedProducts = append(updatedProducts, product)
			}
		} else {
			newProducts = append(newProducts, product)
		}
	}

	if len(newProducts) > 0 {
		if err := r.CreateBatch(ctx, newProducts); err != nil {
			return err
		}
	}

	if len(updatedProducts) > 0 {
		if err := r.UpdateBatch(ctx, updatedProducts); err != nil {
			return err
		}
	}

	return nil
}

func (r *productRepo) GetByArticles(ctx context.Context, articles []string) (map[string]domain.Product, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, article, name, price, image, quantity, category_id FROM products WHERE article IN ($1)",
		pq.Array(articles),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make(map[string]domain.Product)
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Article, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products[product.Article] = product
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) CreateBatch(ctx context.Context, products []domain.Product) error {
	if len(products) == 0 {
    return errors.New("empty products array")
	}
	query := "INSERT INTO products (article, name, price, category_id) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, product := range products {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		values = append(values, product.Article, product.Name, product.Price, product.CategoryID)
	}

	query += strings.Join(placeholders, ", ")
	_, err := r.db.ExecContext(ctx, query, values...)
	return err
}

func (r *productRepo) UpdateBatch(ctx context.Context, products []domain.Product) error {
	if len(products) == 0 {
    return errors.New("empty products array")
	}
	query := "UPDATE products SET name = $1, price = $2, category_id = $3 WHERE id = $4"
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, product := range products {
		_, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.CategoryID, product.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
