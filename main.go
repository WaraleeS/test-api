package main

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Article struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   int       `json:"user_id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
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

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "halo gorgeous"})
		// return c.String(http.StatusOK, fmt.Sprint("halo gorgeous :)"))
	})

	e.GET("/articles", func(c echo.Context) error {
		listArticle := []Article{}
		var a Article

		statement := "SELECT id, title, content, user_id, create_at, update_at delete_at FROM articles"
		results, err := db.Query(statement)
		if err != nil {
			return err
		}
		for results.Next() {
			results.Scan(&a.ID, &a.Title, &a.Content, &a.UserID, &a.CreateAt, &a.UpdateAt, &a.DeleteAt)
			listArticle = append(listArticle, a)
		}

		// for _, val := range articles {
		// 	listArticle = append(listArticle, val)
		// }
		return c.JSON(http.StatusOK, listArticle)
	})

	// e.GET("/articles/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	intID, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
	// 	}
	// 	article, true := articles[intID]
	// 	if !true {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
	// 	}
	// 	return c.JSON(http.StatusOK, article)
	// })

	// e.POST("/articles", func(c echo.Context) error {
	// 	article := Article{}
	// 	max := 0
	// 	err := c.Bind(&article)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	for key := range articles {
	// 		if key > max {
	// 			max = key
	// 		}
	// 	}
	// 	nextKey := max + 1
	// 	article.ID = nextKey
	// 	articles[nextKey] = article

	// 	return c.JSON(http.StatusOK, article)
	// })
	// e.POST("/articles/:id", func(c echo.Context) error {
	// 	article := Article{}
	// 	err := c.Bind(&article)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	articles[article.ID] = article
	// 	return c.JSON(http.StatusOK, article)
	// })
	// e.DELETE("/articles/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	intID, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "sorry girl, Error Response"})
	// 	}
	// 	delete(articles, intID)
	// 	return c.JSON(http.StatusOK, map[string]interface{}{"message": "article deleted!"})
	// })

	e.Logger.Fatal(e.Start(":1234"))
}
