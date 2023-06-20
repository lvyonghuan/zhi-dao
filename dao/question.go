package dao

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"zhi-dao/model"
)

func CreateQuestion(question model.Question) (id int, err error) {
	err = DB.Create(&question).Error
	if err != nil {
		return 0, errors.New("数据库错误，问题创建失败：" + err.Error())
	}
	return question.Id, nil
}

func CreateAnswer(answer model.Answer) (id int, err error) {
	err = DB.Create(&answer).Error
	if err != nil {
		return 0, errors.New("数据库错误，回答创建失败：" + err.Error())
	}
	return answer.Id, nil
}

func SearchQuestionByQuestionID(id int) (question model.Question, err error) {
	err = DB.Where("id=?", id).First(&question).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Question{}, nil
	} else if err != nil {
		return model.Question{}, errors.New("查找问题错误：" + err.Error())
	}
	return question, nil
}

func SearchAnswerByAnswerID(id int) (answer model.Answer, err error) {
	err = DB.Where("id=?", id).First(&answer).Error
	if err != nil {
		return model.Answer{}, errors.New("查找回答错误：" + err.Error())
	}
	return answer, nil
}

func SearchQuestionListByUserID(id int) (questionList model.QuestionList, err error) {
	err = DB.Where("questioner_id=?", id).Find(&questionList).Error
	if err != nil {
		return nil, errors.New("查找用户问题错误：" + err.Error())
	}
	return questionList, nil
}

func SearchAnswerListByUserID(id int) (answerList model.AnswerList, err error) {
	err = DB.Where("answerer_id=?", id).Find(&answerList).Error
	if err != nil {
		return nil, errors.New("查找用户回答错误：" + err.Error())
	}
	return answerList, nil
}

func ChangeQuestion(question model.Question) (err error) {
	err = DB.Save(question).Error
	if err != nil {
		return errors.New("更新问题失败：" + err.Error())
	}
	return nil
}

func ChangeAnswer(answer model.Answer) (err error) {
	err = DB.Save(answer).Error
	if err != nil {
		return errors.New("更新回答失败：" + err.Error())
	}
	return nil
}

func ListAnswerByQuestionID(questionID int) (answerList model.AnswerList, err error) {
	err = DB.Where("question_id=?", questionID).Find(&answerList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, errors.New("查找问题下回答错误：" + err.Error())
	}
	return answerList, nil
}

func DeleteQuestion(question model.Question) (err error) {
	err = DB.Delete(question).Error
	if err != nil {
		return errors.New("删除问题错误：" + err.Error())
	}
	return nil
}

func DeleteAnswer(answer model.Answer) (err error) {
	err = DB.Delete(answer).Error
	if err != nil {
		return errors.New("删除回答错误：" + err.Error())
	}
	return nil
}

func LikeAnswer(answerID, uid int) (err error) {
	isLike, err := Redis.SIsMember(strconv.Itoa(answerID), uid).Result()
	if err != nil {
		return errors.New("redis查询错误：" + err.Error())
	}
	if !isLike {
		err = Redis.SAdd(strconv.Itoa(answerID), uid).Err()
		if err != nil {
			return errors.New("点赞失败")
		}
	} else {
		err = Redis.SRem(strconv.Itoa(answerID), uid).Err()
		if err != nil {
			return errors.New("取消点赞失败")
		}
	}
	err = updateLikes(answerID)
	if err != nil {
		return err
	}
	return nil
}

// 没提前计划好。下次（如果有下次的话，，，）用消息队列试试。
func updateLikes(answerID int) (err error) {
	num, err := Redis.SCard(strconv.Itoa(answerID)).Result()
	if err != nil {
		return errors.New("获取点赞数失败")
	}
	var answer model.Answer
	err = DB.Where("id=?", answerID).First(&answer).Error
	if err != nil {
		return errors.New("查询回答错误")
	}
	answer.Like = int(num)
	err = DB.Save(answer).Error
	if err != nil {
		return errors.New("更新点赞错误")
	}
	return nil
}
