package repositories

import (
	"golang2bookst/internals/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}

// Eksekusi query bisa menggunakan QueryRow
// row := item.db.QueryRow(query, id)

func (item *BookRepo) FindAll() ([]models.BookModel, error) {
	query := "SELECT * FROM books"
	result := []models.BookModel{}
	// .select mengeksekusi query SELECT dan menyimpan hasilnya dalam &result
	if err := item.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (item *BookRepo) FindById(id int) ([]models.BookModel, error) {
	query := "SELECT * FROM books WHERE Id = ?"
	result := []models.BookModel{}
	// .Select mengeksekusi query SELECT dan menyimpan hasilnya dalam &result
	if err := item.Select(&result, query, id); err != nil {
		return nil, err
	}
	return result, nil
}

func (item *BookRepo) DeleteById(id int) error {
	query := "DELETE FROM books WHERE Id = ?"
	// .exec --> tidak mereturn rows (data dari database)
	if _, err := item.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (item *BookRepo) SaveBook(body models.BookModel) error {
	query := "INSERT INTO books(title, description, author) VALUES (?, ?, ?)"
	// .exec --> tidak mereturn rows (data dari database)
	if _, err := item.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil
}

func (item *BookRepo) UpdateById(id int, body models.BookModel) error {
	query := "UPDATE books SET"

	// interface{} dapat menyimpan nilai dari tipe data apa pun
	var args []interface{}

	if body.Title != "" {
		query += " title = ?,"
		args = append(args, body.Title)
	}
	if body.Description != nil && *body.Description != "" {
		query += " description = ?,"
		args = append(args, *body.Description)
	}
	if body.Author != "" {
		query += " author = ?,"
		args = append(args, body.Author)
	}

	// Hapus koma terakhir
	query = strings.TrimSuffix(query, ",")

	query += " WHERE id = ?"
	args = append(args, id)

	// Eksekusi parameterized query dengan Exec, .exec --> tidak mereturn rows (data dari database)
	if _, err := item.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
