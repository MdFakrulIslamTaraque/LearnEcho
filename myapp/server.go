package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GetAdd struct {
	A   int64 `json:"num1"`
	B   int64 `json:"num2"`
	Sum int64 `json:"sum of num1 and num2"`
}

type Person struct {
	Name string `json:"name" form:"name"`
	Age  int64  `json:"age" form:"Age"`
}

type PersonDTO struct {
	Name    string
	Age     int64
	IsAdmin bool
}

// e.GET("/users/:id", getUser)
func getSeum(c echo.Context) error {
	// User ID from path `users/:id`
	id1 := c.Param("a")
	id2 := c.Param("b")
	num1, _ := strconv.ParseInt(id1, 10, 64)
	num2, _ := strconv.ParseInt(id2, 10, 64)
	res := num1 + num2
	sumSt1 := GetAdd{A: num1, B: num2, Sum: res}
	fmt.Println("sum : ", sumSt1)
	// ret := strconv.Itoa(int(res))
	return c.JSON(http.StatusOK, sumSt1)
}

// e.POST("/users/:a/:b", postSum)
// c.Param("id") -- from URL--> "functin-name/:id" format
func postSum(c echo.Context) error {
	// User ID from path `users/:id`
	id1 := c.Param("a")
	id2 := c.Param("b")
	num1, _ := strconv.ParseInt(id1, 10, 64)
	num2, _ := strconv.ParseInt(id2, 10, 64)
	res := num1 - num2
	sumSt1 := GetAdd{A: num1, B: num2, Sum: res}
	// fmt.Println("sum : ", sumSt1)
	// ret := strconv.Itoa(int(res))
	return c.JSON(http.StatusOK, sumSt1)
}

// query Param(get)
// querying the values from the URL after 'show?'(in that part of the URL, we will have 'key'='value' formatted data separated by '&&')
// eg: http://localhost:1323/show?team=bangladesh&&member=taskin&&id=12
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	id := c.QueryParam("id")
	return c.String(http.StatusOK, "team: "+team+", member: "+member+", id: "+id)
}

// e.POST("/save", save)
// FormValue("form-value") -- we will get this value from the form
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// e.POST("/CreatePerson", CreatePerson)
// Bind into struct
func CreatePerson(c echo.Context) error {
	P1 := new(Person)
	if err := c.Bind(P1); err != nil { //by Bind(), the value we are getting from the POST request will be automatically set to the fields of the struct
		return c.String(http.StatusBadRequest, "bad request")
	}

	//updating a new struct, from the updated struct
	U := PersonDTO{
		Name:    P1.Name,
		Age:     P1.Age,
		IsAdmin: false,
	}
	fmt.Println("name: ", U.Name, "age: ", U.Age)
	return c.JSON(http.StatusCreated, U)
}

// POST /upload
// upload the file in the POSTMAN form of body and upload the file there.
func upload(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar") //c.FormFile() will collect the uploaded file through the Key name(creating the handler)
	if err != nil {
		return err
	}
	src, err := avatar.Open() //opening the file through the handler and collect the src address
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(avatar.Filename) //creating the new file and creating a specific address in the local directory
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil { //copying the remote address to the local directory
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thanks, "+name+"!</b>")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:a/:b", getSeum)
	e.POST("/users/:a/:b", postSum)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/upload", upload)
	e.POST("/CreatePerson", CreatePerson)
	e.Logger.Fatal(e.Start(":1323"))
}
