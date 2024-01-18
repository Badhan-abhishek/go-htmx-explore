package main

import (
	"com.quizApp/models"
)

func addSingleQuestionType() {
	println("Adding single question type")
	status := models.QuizType{Label: "Single Question"}
	result := models.DB.Create(&status)
	if result.Error != nil {
		println("Error adding single question type: ", result.Error)
		return
	}
	println("Added single question type")
}

func addMultipleQuestionType() {
	println("Adding multiple question type")
	status := models.QuizType{Label: "Multiple Question"}
	result := models.DB.Create(&status)
	if result.Error != nil {
		println("Error adding multiple question type: ", result.Error)
		return
	}
	println("Added multiple question type")
}

func main() {
	models.ConnectDatabase()
	addSingleQuestionType()
	addMultipleQuestionType()
}
