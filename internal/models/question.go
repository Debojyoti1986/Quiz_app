package models

type Question struct {
	Question 				string		`json:"question"`
	Options 				[]string	`json:"options"`
	CorrectAnswer 			string		`json:"correct_answer"`
	Answer 					string 		`json:"answer"`
}