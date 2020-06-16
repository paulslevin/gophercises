package types

import (
	"fmt"
	"time"
)

type Game struct {
	correctAnswers int
	questions      int
	timeLimit      int
	answerChannel  chan string
}

func NewGame(timeLimit int) Game {
	return Game{
		timeLimit:     timeLimit,
		answerChannel: make(chan string),
	}
}

func (g Game) ParseCSVLines(lines [][]string) {

	timer := g.startTimer()

	for _, line := range lines {

		question := line[0]
		answer := line[1]

		go g.queryUser(question)

		select {
		case <-timer.C:
			g.PrintResult()
			return
		case userAnswer := <-g.answerChannel:
			g.incrementQuestions()
			if userAnswer == answer {
				g.incrementCorrectAnswers()
			}
		}
	}

	g.PrintResult()
}

func (g Game) startTimer() time.Timer {
	timer := time.NewTimer(time.Duration(g.timeLimit) * time.Second)
	return *timer
}

func (g Game) queryUser(question string) {
	var userAnswer string
	fmt.Println("Question:", question)
	fmt.Scan(&userAnswer)
	g.answerChannel <- userAnswer
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
