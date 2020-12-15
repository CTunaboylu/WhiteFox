package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func create_User(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func get_User(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func update_User(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func delete_User(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", create_User)
	e.GET("/users/:id", get_User)
	e.PUT("users/:id", update_User)
	e.DELETE("users/:id", delete_User)

	e.Logger.Fatal(e.Start(":1323"))

	// hello_Handler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello World!\n")
	// }
	// http.HandleFunc("/hello", hello_Handler)
	// log.Println("Listening to requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))
}
