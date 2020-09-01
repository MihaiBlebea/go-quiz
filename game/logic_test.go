package game

import (
	"testing"

	"github.com/MihaiBlebea/go-quiz/player"
	"github.com/MihaiBlebea/go-quiz/quiz"
)

func buildQuiz() (*quiz.Quiz, error) {
	return quiz.New("./../problems.yaml")
}

func TestGameCanBeWon(t *testing.T) {
	quiz, err := buildQuiz()
	if err != nil {
		t.Error(err)
		return
	}

	player := player.Computer{Answers: []string{"London", "15"}}

	gameService := New(10, *quiz, &player)
	score, err := gameService.Run()
	if err != nil {
		t.Error(err)
		return
	}

	if score != 2 {
		t.Errorf("Expected score %d, got score %d", 2, score)
	}
}
