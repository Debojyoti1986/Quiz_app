package main

import (
	"fmt"
	"quizapp/internal/models"
	"quizapp/internal/quiz"
)

func main() {
	var (
		noOfWrongAnswers 		int
		noOfUnanswered 			int
		noOfCorrectAnswered 		int
	)

	core := quiz.NewCore()
	initializeQuestionSet(core)

	for _, question := range core.GetDatastore().Questions {
		core.AnswerQuestion(question, 10)
	}

	fmt.Println("---------------------------")

	for _, question := range core.GetDatastore().Questions {
		if question.CorrectAnswer == question.Answer {
			noOfCorrectAnswered++
		}
	}
	fmt.Println("Number of correct answers: ", noOfCorrectAnswered)

	for _, question := range core.GetDatastore().Questions {
		if question.Answer == "" {
			noOfUnanswered++
		}
	}
	fmt.Println("Number of unanswered questions: ", noOfUnanswered)

	for _, question := range core.GetDatastore().Questions {
		if question.Answer != "" && question.CorrectAnswer != question.Answer {
			noOfWrongAnswers++
		}
	}
	fmt.Println("Number of wrong answers: ", noOfWrongAnswers)
}

func initializeQuestionSet(core quiz.ICore)  {
	question1 := &models.Question{
		Question: "Which year India got independence",
		Options: []string{"1948", "1949", "1947", "1950"},
		CorrectAnswer: "1947",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question1)

	question2 := &models.Question{
		Question: "What is our nation animal",
		Options: []string{"Tiger", "Lion", "Elephant", "Deer"},
		CorrectAnswer: "Tiger",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question2)

	question3 := &models.Question{
		Question: "What is our national sport",
		Options: []string{"Hockey", "Cricket", "Football", "Badminton"},
		CorrectAnswer: "Hockey",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question3)

	question4 := &models.Question{
		Question: "How many states India have",
		Options: []string{"29", "26", "30", "27"},
		CorrectAnswer: "29",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question4)

	question5 := &models.Question{
		Question: "What is our national bird",
		Options: []string{"Eagle", "Sparrow", "Peacock", "Crow"},
		CorrectAnswer: "Peacock",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question5)

	question6 := &models.Question{
		Question: "What is the value of pi",
		Options: []string{"4.12", "3.14", "3.11", "3.26"},
		CorrectAnswer: "3.14",
	}
	core.GetDatastore().Questions = append(core.GetDatastore().Questions, question6)
}
