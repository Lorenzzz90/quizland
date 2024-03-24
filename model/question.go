package model

type QuestionStruct struct {
	Question    string   `json:"question"`
	Answers     []string `json:"answers"`
	Correct     int      `json:"correct"`
	Explanation string   `json:"explanation"`
}
