package models

import (
	"database/sql"
	"fmt"
)

type ArticleRepo struct {
	db *sql.DB
}

// Article Repository
func NewArticleRepository(db *sql.DB) ArticleRepo {
	return ArticleRepo{db: db}
}

// get all database
func (a ArticleRepo) GetAll() ([]Article, error) {
	listArticle := []Article{}

	statement := "SELECT id, title, content, user_id, created_at, updated_at, deleted_at FROM articles"
	results, err := a.db.Query(statement)
	if err != nil {
		return listArticle, err
	}

	for results.Next() {
		a := Article{}
		results.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
		listArticle = append(listArticle, a)
	}
	return listArticle, nil
}

// get data by ID
func (a ArticleRepo) GetByID(intID int) (Article, error) {
	Objarticle := Article{}

	statement := "SELECT * FROM articles WHERE id = ?"
	result := a.db.QueryRow(statement, intID)
	result.Scan(&Objarticle.ID, &Objarticle.Title, &Objarticle.Content, &Objarticle.UserID, &Objarticle.CreatedAt, &Objarticle.UpdatedAt, &Objarticle.DeletedAt)
	return Objarticle, nil
}

// post article to database
func (a ArticleRepo) PostArticle() (Article, error) {
	Objarticle := Article{}

	statement, err := a.db.Prepare("INSERT INTO articles (title, content, user_id) VALUES(?, ?, ?)")
	if err != nil {
		return Objarticle, err
	}
	res, err := statement.Exec(Objarticle.Title, Objarticle.Content, Objarticle.UserID)
	if err != nil {
		return Objarticle, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return Objarticle, err
	}
	Objarticle.ID = int(lastID)
	return Objarticle, nil
}

//	Post article by ID
func (a ArticleRepo) PostArticleByID(id string) (Article, error) {
	Objarticle := Article{}

	statement, err := a.db.Prepare("UPDATE articles SET title=?, content=?, user_id=? WHERE id=?")
	if err != nil {
		return Objarticle, err
	}
	_, err = statement.Exec(Objarticle.Title, Objarticle.Content, Objarticle.UserID, id)
	if err != nil {
		fmt.Print(err)
		return Objarticle, err
	}
	return Objarticle, nil
}

// update delete_at from database
func (a ArticleRepo) DeletedArticle(id string) (string, error) {
	messageError := "Error delete database"
	massageSuccess := "Delete Successful !!"

	statement, err := a.db.Prepare("UPDATE articles SET deleted_at=NOW() where id=?")
	if err != nil {
		return messageError, err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return messageError, err
	}
	return massageSuccess, nil
}
