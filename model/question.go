package model

type Question struct {
	Id           int    `json:"id"`
	QuestionerId int    `json:"questioner_id"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"` //问题介绍
	Topic        string `json:"topic"`     //分类
}
