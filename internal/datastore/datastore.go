package datastore

import (
	"quizapp/internal/models"
)

type Datastore struct {
	Questions []*models.Question
}

type IDatastore interface {
	GetQuestionList() 	[]models.Question
}

func (d Datastore) GetQuestionMap() []*models.Question {
	return d.Questions
}
