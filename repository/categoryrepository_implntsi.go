package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_resfullapi/helper"
	"golang_resfullapi/model/domain"
)

type CategoryRepositoryImplementasi struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImplementasi{}
}

func (repository *CategoryRepositoryImplementasi) Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category, nil
}

func (repository *CategoryRepositoryImplementasi) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImplementasi) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImplementasi) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId) // Menggunakan parameter yang benar
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return domain.Category{}, errors.New("category tidak ada")
	}
}

func (repository *CategoryRepositoryImplementasi) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "Select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
