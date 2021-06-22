package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"quizapp/internal/datastore"
	"quizapp/internal/models"
	"strings"
	"sync"
	"time"
)

var (
	once 	sync.Once
)

type Core struct {
	Datastore *datastore.Datastore
}

type ICore interface {
	AddQuestion(questionText string, options []string, correctAnswer string)
	AnswerQuestion(question *models.Question, timeout time.Duration) error
	GetDatastore() *datastore.Datastore
}

func NewCore() ICore {
	var datastoreInstance *datastore.Datastore
	once.Do(func() {
		datastoreInstance = &datastore.Datastore{
			Questions: []*models.Question{},
		}
	})
	return &Core{
		Datastore: datastoreInstance,
	}
}

func (c *Core) AddQuestion(questionText string, options []string, correctAnswer string) {
	question := &models.Question{
		Question: questionText,
		Options: options,
		CorrectAnswer: correctAnswer,
		Answer: "",
	}
	c.Datastore.Questions = append(c.Datastore.Questions, question)
}

func (c *Core) AnswerQuestion(question *models.Question, timeout time.Duration) error {
	input := make(chan string, 1)

	fmt.Println("---------------------------")
	fmt.Println(question.Question)
	for idx, option := range question.Options {
		fmt.Println(idx+1, ": " + option)
	}

	fmt.Println("Please choose from the options")

	go getAnswer(input)

	select {
	case answer := <-input:
		answer = strings.Replace(answer, "\n", "", -1)
		question.Answer = answer
		if question.CorrectAnswer == answer {
			fmt.Println("Correct answer")
		} else {
			fmt.Println("Wrong answer. Correct answer is ", question.CorrectAnswer)
		}
	case <-time.After(timeout * time.Second):
		fmt.Println("Time out")
	}

	return nil
}

func (c *Core)GetDatastore() *datastore.Datastore {
	return c.Datastore
}

func closeChannel(input chan string)  {
	close(input)
}

func getAnswer(input chan string) {
	defer closeChannel(input)
	in := bufio.NewReader(os.Stdin)
	result, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input <- result
}
