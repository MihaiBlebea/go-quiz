package game

import (
	"testing"
	"time"

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

func TestGameExpires(t *testing.T) {
	quiz, err := buildQuiz()
	if err != nil {
		t.Error(err)
		return
	}

	player := player.Computer{Answers: []string{"London", "15"}}

	start := time.Now()
	gameService := New(0, *quiz, &player)
	score, err := gameService.Run()
	if err != nil {
		t.Error(err)
		return
	}

	duration := int(time.Since(start).Seconds())

	if duration > 0 {
		t.Errorf("Expected duration to be %d, got score %d", 0, duration)
	}

	if score == 2 {
		t.Errorf("Expected score to not be %d, got score %d", 2, score)
	}
}
