package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"strconv"

	"github.com/skratchdot/open-golang/open"
)

var (
	settings saveData
)

func main() {
	listen := make(chan bool)
	go openBrowser(listen)

	e := echo.New()
	e.Static("/", "assets")
	e.GET("/save", execSave)
	listen <- true
	e.Logger.Fatal(e.Start(":1323"))
}

func openBrowser(listen chan bool) {
	<-listen
	time.Sleep(100 * time.Millisecond)
	open.Run("http://localhost:1323")
}

func execSave(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func init() {
	cacheTypes := []int{1, 5, 10, 50, 100, 500, 1000, 5000, 10000}
	for _, cacheType := range cacheTypes {
		settings.caches = append(settings.caches, cache{50, cacheType, strconv.Itoa(cacheType) + "yen", 0})
	}
}

type saveData struct {
	date          string
	caches        []cache
	sales         int
	otherServices []otherService
	unpaid        []int
	in            []int
	out           []int
	others        []int
	diff          int
}

type otherService struct {
	unitCost   int
	n          int
	isPositive int
	name       string
}

type cache struct {
	initialValue int
	unitCost     int
	name         string
	value        int
}
