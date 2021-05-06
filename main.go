package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
	"log"
	// "encoding/csv"
	"time"
	"strconv"
)

var (
	settings saveData
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

func init() {
	cacheTypes := []int {1, 5, 10, 50, 100, 500, 1000, 5000, 10000}
	for _, cacheType := range cacheTypes {
		settings.caches = append(settings.caches, cache{50, cacheType, strconv.Itoa(cacheType) + "yen", 0})
	}
}

type saveData struct {
	date string
	caches []cache
	sales int
	otherServices []otherService
	unpaid []int
	in []int
	out []int
	others []int
	diff int
}

type otherService struct {
	unitCost int
	n int
	isPositive int
	name string
}

type cache struct {
	initialValue int
	unitCost int
	name string
	value int
}
