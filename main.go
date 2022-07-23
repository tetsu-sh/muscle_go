package main

import (
	"github.com/labstack/echo/v4"
	"github.com/muscle_go/db"
	"github.com/muscle_go/presentation"
	"github.com/muscle_go/repository"
)

func main() {
	e := echo.New()

	d := db.NewDB()
	tr := repository.TrainRepository(d)
	h := presentation.NewHandler(tr)
	h.Register(&e)
	e.Logger.Fatal(e.Start(":8000"))
}
