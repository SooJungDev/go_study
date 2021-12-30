package main

import (
	"github.com/labstack/echo"
	"goForBeginner/scrapper"
	"os"
	"strings"
)

const fileName string = "jobs.csv"

// Handler
func handleHome(c echo.Context) error {
	return c.File("/Users/crytal2840/go/src/go_study/src/goForBeginner/home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
