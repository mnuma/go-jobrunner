package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	fmt.Println("start")
	jobrunner.Start(10)
	jobrunner.Schedule("0 0-2,6-23 * * *", Task{})
	jobrunner.Schedule("@every 1s", Task{})

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, jobrunner.StatusJson())
	})

	e.Start(":1231")
}

type Task struct {
}

func (e Task) Run() {
	fmt.Printf("job..... \n")
}

