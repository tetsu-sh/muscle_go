package main

import (
	"muscle_go/cmd/db"

	"muscle_go/cmd/presentation"
	"muscle_go/cmd/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	d := db.NewDB()
	tr := repository.NewTrainRepository(d)
	h := presentation.NewHandler(tr)
	h.Register(e)
	e.Logger.Fatal(e.Start(":8000"))
}
