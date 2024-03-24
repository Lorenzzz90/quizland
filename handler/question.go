package handler

import (
	"github.com/lorenzzz90/quizland/model"
	"github.com/lorenzzz90/quizland/view/question"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	Questions []model.QuestionStruct
	Index     int
	Score     int
}

func (h *QuestionHandler) HandleQuestionShow(c echo.Context) error {
	h.Index = 0
	h.Score = 0
	return render(c, question.Show(h.Questions[h.Index]))
}

func (h *QuestionHandler) HandleQuestionPost(c echo.Context) error {
	c.Request().ParseForm()
	answer := c.Request().FormValue("answer")
	intAnswer, err := strconv.Atoi(answer)
	if err != nil {
		http.Error(c.Response().Writer, "Invalid option", http.StatusTeapot)
		return err
	}
	if intAnswer == h.Questions[h.Index].Correct {
		h.Score++
	}
	if h.Index+1 < len(h.Questions) {
		h.Index++
		return render(c, question.Show(h.Questions[h.Index]))
	} else {
		score := h.Score
		h.Index = 0
		h.Score = 0
		return render(c, question.Result(score))
	}
}
