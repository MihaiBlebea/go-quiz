package quiz

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Quiz _
type Quiz struct {
	Problems []Problem `yaml:"problems"`
}

// Problem _
type Problem struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

// New returns a new Quiz
func New(fileName string) (*Quiz, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &Quiz{}, err
	}

	var quiz Quiz
	err = yaml.Unmarshal(b, &quiz)
	if err != nil {
		return &Quiz{}, err
	}

	return &quiz, nil
}
