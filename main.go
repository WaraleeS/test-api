package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User_ID string `json:"user_id"`
}

type User struct {
	User_ID  string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	articles := map[int]Article{
		1: {
			ID:      1,
			Title:   "Thailand",
			Content: "this is thailand content",
			User_ID: "11111",
		},
		2: {
			ID:      2,
			Title:   "US",
			Content: "this is US content",
			User_ID: "222222",
		},
		3: {
			ID:      3,
			Title:   "Germany",
			Content: "this is Germany content",
			User_ID: "33333",
		},
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "halo gorgeous"})
		// return c.String(http.StatusOK, fmt.Sprint("halo gorgeous :)"))
	})

	e.GET("/articles", func(c echo.Context) error {
		// article := &Article{
		// 	ID:      1,
		// 	Title:   "Annyoeng@Haseyo.com",
		// 	Content: "Covid-19",
		// 	User_ID: "12345678",
		// }
		// return c.JSON(http.StatusOK, article)

		listArticle := []Article{}
		for _, val := range articles {
			listArticle = append(listArticle, val)
		}
		return c.JSON(http.StatusOK, listArticle)

	})

	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{"message": "sorry girl, Error Response"})
		}
		listArticle, true := articles[intID]
		if !true {
			return c.JSON(http.StatusOK, map[string]interface{}{"message": "sorry girl, Error Response"})
		}
		return c.JSON(http.StatusOK, listArticle)
	})

	e.POST("/articles", func(c echo.Context) error {
		article := Article{}
		maxKeyID := 0
		err := c.Bind(&article)
		if err != nil {
			return err
		}
		for key := range articles {
			if key > maxKeyID {
				maxKeyID = key
			}
		}
		nextKeyID := maxKeyID + 1
		article.ID = nextKeyID
		articles[nextKeyID] = article

		article = Article{
			ID:      nextKeyID,
			Title:   "Italy",
			Content: "This is Italy content",
			User_ID: "44444",
		}
		return c.JSON(http.StatusOK, article)
	})
	e.POST("/articles/:id", func(c echo.Context) error {
		article := Article{}
		err := c.Bind(&article)
		if err != nil {
			return err
		}
		articles[article.ID] = article
		return c.JSON(http.StatusOK, article)
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("update article"))
	})

	e.Logger.Fatal(e.Start(":1234"))
}
