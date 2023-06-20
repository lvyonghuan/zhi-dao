package model

type CommentList []Comment

type Comment struct {
	Id          int    `json:"id"`
	QuestionId  int    `json:"question_id"`
	AnswerId    int    `json:"answer_id"`
	CommenterId int    `json:"commenter_id"`
	ReplyId     int    `json:"reply_id"`
	Text        string `json:"text"`
	Like        int    `json:"like"`
}
