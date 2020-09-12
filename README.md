# Go Quiz Game

Objectives
- how to start a project in GO?
- how to use dependencies in GO?
- do we always need the a main.go file? 
- how do structs compare with classes in PHP?
- how can we define an interface and implement it?

## Create a new GO project
- Init the module `go mod init github.com/MihaiBlebea/go-quiz`

## Add flags for questions file and timeout
- Add 2 flags:
    - Flag for adding the path to the problems file - default to problems.yaml
    - Flag for adding the timeout - default to 10 sec

For adding a flag:
```go
fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
```

Do not forget to call parse after definiting the flags:
```go
flag.Parse()
```

## Read questions file and create structs
Add a dependency package to decode the yaml to struct:
- `gopkg.in/yaml.v2`

Quiz struct:
```go
type Quiz struct {
	Problems []Problem `yaml:"problems"`
}
```

Problem struct:
```go
type Problem struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}
```

Read the file and parse to Quiz struct
```go
b, err := ioutil.ReadFile(fileName)
if err != nil {
    return log.Fatal(err)
}

var quiz Quiz
err = yaml.Unmarshal(b, &quiz)
if err != nil {
    return log.Fatal(err)
}
```

## Create the game loop

### 1. First version
```go
var score int
for i, prob := range quiz.Problems {
    fmt.Println(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))

    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    
    if input == prob.Answer {
        fmt.Println("Correct")
        score++
    } else {
        fmt.Println("Wrong")
    }
}

fmt.Println(fmt.Sprintf("Game over! Your score is %d from %d", score, len(quiz.Problems)))
```

### 2. Add timeout timer
- Add timer outside of the for loop:
```go
t := time.NewTimer(time.Duration(s.limit) * time.Second)
```

- Add select to check if the timer has expired before the input check
```go
var score int
for i, prob := range quiz.Problems {
    fmt.Println(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))

    select {
    case <-t.C:
        break
    default:
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        
        if input == prob.Answer {
            fmt.Println("Correct")
            score++
        } else {
            fmt.Println("Wrong")
        }
    }
}

fmt.Println(fmt.Sprintf("Game over! Your score is %d from %d", score, len(quiz.Problems)))
```

### 3. Add channel to read the answers
```go
var score int
for i, prob := range quiz.Problems {
    fmt.Println(fmt.Sprintf("Problem %d: %s ?", i+1, prob.Question))

    // Create answer as type for the channel
    type Answer struct {
        input string
        err   error
    }

    // Create a channel of type Answer
    answerChan := make(chan Answer)

    // Create a routine to get the user input
    go func() {
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')

        if err != nil {
            answerChan <- Answer{
                err: err,
            }
        }

        // Trim the space on the input
        input = strings.TrimSpace(input)

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
            fmt.Println("Correct")
            score++
        } else {
            fmt.Println("Wrong")
        }
    }
}

fmt.Println(fmt.Sprintf("Game over! Your score is %d from %d", score, len(quiz.Problems)))
```

## Improve: Add quiz package
Points to touch:
- What is the different between a struct and a package
- What is the different between a method with a receiver and just a function
- How does the constructor works in go

```go
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
```

Create a simple test for the quiz package
```go
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
```

## Improve: Add game package
```go
// Player interface. Replace printing ot the console with the Print method on the player
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
	// Put all the logic in here

	return score, nil
}
```

## Improve: Create a player package
Human player
```go
// Human player
type Human struct {
}

// Print _
func (h *Human) Print(output string) {
	fmt.Println(output)
}

// Input _
func (h *Human) Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
```

Computer player
```go
// Computer player
type Computer struct {
	Answers []string
	counter int
}

// Print _
func (c *Computer) Print(output string) {
	fmt.Println(output)
}

// Input _
func (c *Computer) Input() (string, error) {
	defer c.incrementCounter()

	return c.Answers[c.counter], nil
}

func (c *Computer) incrementCounter() {
	c.counter++
}
```

## Puting it all back together
```go
func main() {
	fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
	limit := flag.Int("limit", 10, "The time limit for the quiz in seconds")
	flag.Parse()

	quiz, err := quiz.New(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	// player := player.Human{}
	player := player.Computer{Answers: []string{"London", "15"}}

	gameService := game.New(*limit, *quiz, &player)
	_, err = gameService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
```