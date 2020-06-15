package types

import (
	"fmt"
	"time"
)

type Game struct {
	correctAnswers int
	questions      int
	TimeLimit      int
}

func (g *Game) ParseCSVLines(lines [][]string) {

	timer := g.startTimer()

	for _, line := range lines {

		question := line[0]
		answer := line[1]

		answerChannel := make(chan string)

		go queryUser(question, answerChannel)

		select {
		case <-timer.C:
			g.PrintResult()
			return
		case userAnswer := <-answerChannel:
			g.incrementQuestions()
			if userAnswer == answer {
				g.incrementCorrectAnswers()
			}
		}
	}

	g.PrintResult()
}

func (g *Game) startTimer() *time.Timer {
	timer := time.NewTimer(time.Duration(g.TimeLimit) * time.Second)
	return timer
}

func (g *Game) incrementCorrectAnswers() {
	g.correctAnswers++
}

func (g *Game) incrementQuestions() {
	g.questions++
}

func (g Game) PrintResult() {
	fmt.Println("Number of questions answered:", g.questions)
	fmt.Println("Number of correct answers:", g.correctAnswers)
}

func queryUser(question string, answerChannel chan string) {
	var userAnswer string
	fmt.Println("Question:", question)
	fmt.Scan(&userAnswer)
	answerChannel <- userAnswer
}
