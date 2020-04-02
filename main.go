package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func main() {
	// articles := map[int]Article{
	// 	1: {
	// 		ID:      1,
	// 		Title:   "Thailand",
	// 		Content: "this is thailand content",
	// 		UserID:  11111,
	// 	},
	// 	2: {
	// 		ID:      2,
	// 		Title:   "US",
	// 		Content: "this is US content",
	// 		UserID:  222222,
	// 	},
	// 	3: {
	// 		ID:      3,
	// 		Title:   "Germany",
	// 		Content: "this is Germany content",
	// 		UserID:  33333,
	// 	},
	// }

	// connect database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "halo gorgeous"})
		// return c.String(http.StatusOK, fmt.Sprint("halo gorgeous :)"))
	})

	e.GET("/articles", func(c echo.Context) error {
		listArticle := []Article{}

		statement := "SELECT id, title, content, user_id, created_at, updated_at, deleted_at FROM articles"
		results, err := db.Query(statement)
		if err != nil {
			return err
		}
		for results.Next() {
			a := Article{}
			results.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
			listArticle = append(listArticle, a)
		}

		// for _, val := range articles {
		// 	listArticle = append(listArticle, val)
		// }
		return c.JSON(http.StatusOK, listArticle)
	})

	e.GET("/articles/:id", func(c echo.Context) error {
		// id := c.Param("id")
		// intID, err := strconv.Atoi(id)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
		// }
		// article, true := articles[intID]
		// if !true {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
		// }
		// return c.JSON(http.StatusOK, article)

		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Error select database !!"})
		}
		a := Article{}
		statement := "SELECT * FROM articles WHERE id = ?"
		result := db.QueryRow(statement, intID)
		result.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
		return c.JSON(http.StatusOK, a)
	})

	e.POST("/articles", func(c echo.Context) error {
		a := Article{}
		statement, err := db.Prepare("INSERT INTO articles (title, content, user_id) VALUES(?, ?, ?)")
		if err != nil {
			return err
		}
		res, err := statement.Exec(a.Title, a.Content, a.UserID)
		if err != nil {
			return err
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		a.ID = int(lastID)
		return c.JSON(http.StatusOK, a)

		// article := Article{}
		// max := 0
		// err := c.Bind(&article)
		// if err != nil {
		// 	return err
		// }
		// for key := range articles {
		// 	if key > max {
		// 		max = key
		// 	}
		// }
		// nextKey := max + 1
		// article.ID = nextKey
		// articles[nextKey] = article
		// return c.JSON(http.StatusOK, article)
	})

	e.POST("/articles/:id", func(c echo.Context) error {
		a := Article{}
		id := c.Param("id")
		err = c.Bind(&a)
		if err != nil {
			return err
		}
		statement, err := db.Prepare("UPDATE articles SET title=?, content=?, user_id=? WHERE id=?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(a.Title, a.Content, a.UserID, id)
		if err != nil {
			fmt.Print(err)
			return err
		}
		return c.JSON(http.StatusOK, a)

		// article := Article{}
		// err := c.Bind(&article)
		// if err != nil {
		// 	return err
		// }
		// articles[article.ID] = article
		// return c.JSON(http.StatusOK, article)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		// id := c.Param("id")
		// intID, err := strconv.Atoi(id)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
		// }
		// statement := "DELETE FROM articles WHERE id = ?"
		// _, err = db.Exec(statement, intID)
		// if err != nil {
		// 	return err
		// }
		// return c.JSON(http.StatusOK, map[string]interface{}{"message": "article deleted!"})

		id := c.Param("id")
		statement, err := db.Prepare("UPDATE articles SET deleted_at=NOW() where id=?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Delete Successful !!"})

		// id := c.Param("id")
		// intID, err := strconv.Atoi(id)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
		// }
		// delete(articles, intID)
		// return c.JSON(http.StatusOK, map[string]interface{}{"message": "article deleted!"})
	})

	e.Logger.Fatal(e.Start(":1234"))
}
