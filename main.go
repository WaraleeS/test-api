package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/WaraleeS/test-api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// connect database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	if err != nil {
		panic(err)
	}

	articleRepository := models.NewArticleRepository(db)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "halo gorgeous"})
	})

	e.GET("/articles", func(c echo.Context) error {
		ListArticle, err := articleRepository.GetAll()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ListArticle)

		// listArticle := []Article{}

		// statement := "SELECT id, title, content, user_id, created_at, updated_at, deleted_at FROM articles"
		// results, err := db.Query(statement)
		// if err != nil {
		// 	return err
		// }
		// for results.Next() {
		// 	a := Article{}
		// 	results.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
		// 	listArticle = append(listArticle, a)
		// }
	})

	e.GET("/articles/:id", func(c echo.Context) error {
		// id := c.Param("id")
		// intID, err := strconv.Atoi(id)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Error select database !!"})
		// }
		// a := Article{}
		// statement := "SELECT * FROM articles WHERE id = ?"
		// result := db.QueryRow(statement, intID)
		// result.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
		// return c.JSON(http.StatusOK, a)

		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		ListArticle, err := articleRepository.GetByID(intID)
		if err != nil {
			fmt.Print(err)
			return err
		}
		return c.JSON(http.StatusOK, ListArticle)
	})

	e.POST("/articles", func(c echo.Context) error {
		ListArticle, err := articleRepository.PostArticle()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ListArticle)

		// a := Article{}
		// statement, err := db.Prepare("INSERT INTO articles (title, content, user_id) VALUES(?, ?, ?)")
		// if err != nil {
		// 	return err
		// }
		// res, err := statement.Exec(a.Title, a.Content, a.UserID)
		// if err != nil {
		// 	return err
		// }
		// lastID, err := res.LastInsertId()
		// if err != nil {
		// 	return err
		// }
		// a.ID = int(lastID)
		// return c.JSON(http.StatusOK, a)
	})

	e.POST("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		ListArticle, err := articleRepository.PostArticleByID(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ListArticle)

		// a := Article{}
		// id := c.Param("id")
		// err = c.Bind(&a)
		// if err != nil {
		// 	return err
		// }
		// statement, err := db.Prepare("UPDATE articles SET title=?, content=?, user_id=? WHERE id=?")
		// if err != nil {
		// 	return err
		// }
		// _, err = statement.Exec(a.Title, a.Content, a.UserID, id)
		// if err != nil {
		// 	fmt.Print(err)
		// 	return err
		// }
		// return c.JSON(http.StatusOK, a)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		ListArticle, err := articleRepository.DeletedArticle(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ListArticle)

		// id := c.Param("id")
		// statement, err := db.Prepare("UPDATE articles SET deleted_at=NOW() where id=?")
		// if err != nil {
		// 	return err
		// }
		// _, err = statement.Exec(id)
		// if err != nil {
		// 	return err
		// }
		// return c.JSON(http.StatusOK, map[string]interface{}{"message": "Delete Successful !!"})
	})

	e.Logger.Fatal(e.Start(":1234"))
}
