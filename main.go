package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
	"log"
	"encoding/csv"
	"time"
)

func main() {
	e := echo.New()
	e.Static("/", "assets")
	e.GET("/save", save)
	e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := "hello from go"
	file.Write(([]byte)(output))
	return c.String(http.StatusOK, "Hello, World")
}

type savaData struct {
	date time.Time
	yen1, yen5, yen10, yen50, yen100, yen500, yen1000, yen5000, yen10000 int
	sales int
	rabiesContinuation int
	rabiesNew int
	in int
	out int
	others int
	diff int
}

type baseData struct {
	yen1, yen5, yen10, yen50, yen100, yen500, yen1000, yen5000, yen10000 int
	rabiesNewCost int
	rabiesContinuationCost int
}
