package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/Lorenzzz90/quizland/handler"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	questionHandler := handler.QuestionHandler{Index: 0, Score: 0}

	app.GET("/", questionHandler.MainPage)
	app.POST("/next", questionHandler.Next)
	app.POST("/explain", questionHandler.Explain)

	app.Static("/css", "css")
	app.Start(":3000")
	openbrowser("localhost:3000")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
