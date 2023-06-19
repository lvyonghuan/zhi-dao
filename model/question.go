package model

type Question struct {
	Id           int    `json:"id"`
	QuestionerId int    `json:"questioner_id"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"` //问题介绍
	Topic        string `json:"topic"`     //分类
}

type Answer struct {
	Id         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	AnswererId int    `json:"answerer_id"`
	Text       string `json:"text"`
	Like       int    `json:"like"` //点赞数
}
