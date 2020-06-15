package types

import (
	"testing"
)

func TestGame(t *testing.T) {
	game := new(Game)
	correctAnswers := game.correctAnswers
	questions := game.questions
	timeLimit := game.TimeLimit

	if correctAnswers != 0 {
		t.Errorf("Game initialized with correctAnswers=%d, expected 0", correctAnswers)
	} else {
		t.Log("Game initialized successfully with correctAnswers=0")
	}

	if questions != 0 {
		t.Errorf("Game initialized with questions=%d, expected 0", questions)
	} else {
		t.Log("Game initialized successfully with questions=0")
	}

	if timeLimit != 0 {
		t.Errorf("Game initialized with TimeLimit=%d, expected 0", timeLimit)
	} else {
		t.Log("Game initialized successfully with TimeLimit=0")
	}

	for i := 0; i < 20; i++ {
		game.incrementCorrectAnswers()
		if i%2 == 0 {
			game.incrementQuestions()
		}
	}

	correctAnswers = game.correctAnswers
	if correctAnswers != 20 {
		t.Errorf("After incrementing, correctAnswers=%d, expected 20", correctAnswers)
	} else {
		t.Log("After incrementing, correctAnswers=20 as expected")
	}

	questions = game.questions
	if questions != 10 {
		t.Errorf("After incrementing, questions=%d, expected 10", questions)
	} else {
		t.Log("After incrementing, questions=10 as expected")
	}
}
