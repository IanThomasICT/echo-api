package main

import (
	"net/http"

	"github.com/carlmjohnson/requests"
	"github.com/labstack/echo/v4"
)

type CatFact struct {
	Fact   string `json:"fact"`  
	Length int64  `json:"length"`
}

type ResponseObj struct {
	Ok	bool `json:"ok"`	
	Data interface{} `json:"data"`
}



func main() {
	e := echo.New();

	catFactHandler := func(c echo.Context) error {
		var data CatFact
		
		if err := requests.
			URL("https://catfact.ninja/fact").
			ToJSON(&data).
			Fetch(c.Request().Context()); err != nil {
				return c.JSON(http.StatusInternalServerError, new(ResponseObj))
		}
		
		return c.JSON(200, &ResponseObj{true, data})
	}

	e.GET("/cats", catFactHandler)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "it works")
	})
	
	e.Logger.Fatal(e.Start(":8080"))
	
}
