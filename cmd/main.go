package main

import (
	"encoding/json"
	"fmt"
	"github.com/Lorenzzz90/quizland/handler"
	"github.com/Lorenzzz90/quizland/model"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/labstack/echo/v4"
)

func main() {
	var questions []model.QuestionStruct
	file, err := os.Open("goquiz.json")
	if err != nil {
		log.Fatal(err)
	}
	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(read, &questions)
	app := echo.New()
	//app.Use(middleware.Logger())
	questionHandler := handler.QuestionHandler{Questions: questions, Index: 0, Score: 0}

	app.GET("/", questionHandler.HandleQuestionShow)
	app.POST("/", questionHandler.HandleQuestionPost)

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
