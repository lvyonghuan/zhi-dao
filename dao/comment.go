package dao

import (
	"errors"
	"zhi-dao/model"
)

func CreateComment(comment model.Comment) (id int, err error) {
	err = DB.Create(&comment).Error
	if err != nil {
		return 0, errors.New("创建评论错误：" + err.Error())
	}
	return comment.Id, nil
}

func SearchCommentByCommentID(commentID int) (comment model.Comment, err error) {
	err = DB.Where("id=?", commentID).First(&comment).Error
	if err != nil {
		return model.Comment{}, errors.New("查找评论错误：" + err.Error())
	}
	return comment, nil
}

func DeleteComment(comment model.Comment) (err error) {
	err = DB.Delete(comment).Error
	if err != nil {
		return errors.New("删除评论错误：" + err.Error())
	}
	return nil
}

func GetCommentList(answerID int) (commentList model.CommentList, err error) {
	err = DB.Where("answer_id=?", answerID).Find(&commentList).Error
	if err != nil {
		return nil, errors.New("查询评论错误：" + err.Error())
	}
	return commentList, nil
}

func GetReplyCommentList(replyID int) (commentList model.CommentList, err error) {
	err = DB.Where("reply_id=?", replyID).Find(&commentList).Error
	if err != nil {
		return nil, errors.New("查询评论错误：" + err.Error())
	}
	return commentList, nil
}
