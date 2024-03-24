package handler

import (
	"net/http"
	"strconv"

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
	c.Request().ParseForm()
	answer := c.Request().FormValue("answer")
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
