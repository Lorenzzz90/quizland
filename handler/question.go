package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Lorenzzz90/quizland/model"
	"github.com/Lorenzzz90/quizland/view/question"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	Questions []model.QuestionStruct
	Index     int
	Score     int
}

func (h *QuestionHandler) MainPage(c echo.Context) error {
	h.Index = 0
	h.Score = 0
	return render(c, question.Start())
}

func (h *QuestionHandler) Next(c echo.Context) error {
	if h.Index == 0 {
		var questions []model.QuestionStruct
		var file *os.File
		var err error
		files, err := os.ReadDir("./quizFiles")
		if err != nil {
			log.Fatal("Impossibile leggere file nella directory")
		}
		for _, f := range files {
			fmt.Println(fmt.Sprint(strings.ToLower(c.FormValue("quiz")), ".json"))
			if f.Name() == fmt.Sprint(strings.ToLower(c.FormValue("quiz")), ".json") {
				file, err = os.Open(fmt.Sprint("./quizFiles/", f.Name()))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		defer file.Close()
		read, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(read, &questions)
		if err != nil {
			log.Fatal(err)
		}
		h.Questions = questions
	}
	if h.Index < len(h.Questions) {
		h.Index++
		return render(c, question.Show(h.Questions[h.Index-1]))
	} else {
		score := h.Score
		h.Index = 0
		h.Score = 0
		return render(c, question.Result(score, len(h.Questions)))
	}
}

func (h *QuestionHandler) Explain(c echo.Context) error {
	correct := false
	answer := c.FormValue("answer")
	intAnswer, err := strconv.Atoi(answer)
	if err != nil {
		http.Error(c.Response().Writer, "Invalid option", http.StatusTeapot)
		return err
	}
	if intAnswer == h.Questions[h.Index-1].Correct {
		h.Score++
		correct = true
	}
	return render(c, question.Explain(h.Questions[h.Index-1].Explanation, correct))
}
