package quiz

import "testing"

func TestLoadProblemsFromFile(t *testing.T) {
	quiz, err := New("./../problems.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if len(quiz.Problems) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(quiz.Problems))
	}

	if quiz.Problems[0].Question != "What is the capital of UK" {
		t.Errorf("Expected %s, got %s", "What is the capital of UK", quiz.Problems[0].Question)
	}

	if quiz.Problems[0].Answer != "London" {
		t.Errorf("Expected %s, got %s", "London", quiz.Problems[0].Answer)
	}
}
