package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/MihaiBlebea/go-quiz/quiz"
)

// Player interface
type Player interface {
	Print(string)
	Input() (string, error)
}

type service struct {
	limit  int
	quiz   quiz.Quiz
	player Player
}

// New _
func New(limit int, quiz quiz.Quiz, player Player) Service {
	return &service{limit, quiz, player}
}

func (s *service) Run() (int, error) {
	t := time.NewTimer(time.Duration(s.limit) * time.Second)

	var score int
	for i, prob := range s.quiz.Problems {
		s.player.Print(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))

		type Answer struct {
			input string
			err   error
		}

		answerChan := make(chan Answer)

		go func() {
			input, err := s.player.Input()
			if err != nil {
				answerChan <- Answer{
					err: err,
				}
			}

			input = strings.TrimSpace(input)
			if err != nil {
				answerChan <- Answer{
					err: err,
				}
			}

			answerChan <- Answer{
				input: input,
			}
		}()

		select {
		case <-t.C:
			break
		case answer := <-answerChan:
			if answer.err != nil {
				return score, answer.err
			}

			if answer.input == prob.Answer {
				s.player.Print("Correct\n")
				score++
			} else {
				s.player.Print("Wrong\n")
			}
		}
	}

	s.player.Print(fmt.Sprintf("Game over! Your score is %d from %d", score, len(s.quiz.Problems)))

	return score, nil
}
