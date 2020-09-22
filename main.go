package main

import (
	"flag"
	"log"

	"github.com/MihaiBlebea/go-quiz/game"
	"github.com/MihaiBlebea/go-quiz/player"
	"github.com/MihaiBlebea/go-quiz/quiz"
)

func main() {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 10, "The time limit for the quiz in seconds")
	flag.Parse()

	quiz, err := quiz.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	player := player.Human{}
	// player := player.Computer{Answers: []string{"London", "15"}}

	gameService := game.New(*limit, *quiz, &player)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
