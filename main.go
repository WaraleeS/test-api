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
			ID:      11,
			Title:   "Thailand",
			Content: "this is thailand content",
			User_ID: "11111",
		},
		2: {
			ID:      22,
			Title:   "US",
			Content: "this is US content",
			User_ID: "222222",
		},
		3: {
			ID:      33,
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

		list := []Article{}
		for _, val := range articles {
			list = append(list, val)
		}
		return c.JSON(http.StatusOK, list)

	})
	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		newID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		article := &Article{
			ID:      newID,
			Title:   "Annyoeng@Haseyo.com",
			Content: "Covid-19",
			User_ID: "12345678",
		}
		return c.JSON(http.StatusOK, article)
	})
	e.POST("/articles", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("insert article"))
	})
	e.POST("/articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("update article"))
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("update article"))
	})

	e.Logger.Fatal(e.Start(":1234"))
}
