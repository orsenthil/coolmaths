package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

type Main struct {
	styles    *Styles
	index     int
	questions [3]Question
	width     int
	height    int
	done      bool
}

type Question struct {
	question string
	expected int
	answer   string
	input    Input
}

func newQuestion(question string) Question {
	return Question{question: question}
}

func newShortQuestion(q string, expected int) Question {
	question := newQuestion(q)
	model := NewShortAnswerField()
	question.input = model
	question.expected = expected
	return question
}
func newLongQuestion(q string) Question {
	question := newQuestion(q)
	model := NewLongAnswerField()
	question.input = model
	return question
}

func New(questions [3]Question) *Main {
	styles := DefaultStyles()
	return &Main{questions: questions, styles: styles}
}

func (m Main) Init() tea.Cmd {
	return m.questions[m.index].input.Blink
}
func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	current := &m.questions[m.index]
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			}
			current.answer = current.input.Value()
			ans, _ := strconv.Atoi(current.answer)
			var response string
			if ans == current.expected {
				response = "You got it right!"
			} else {
				response = "You were wrong!"
			}

			log.Printf("Question: %s, Your Answer: %s. Correct Answer: %d. %s ", current.question, current.answer, current.expected, response)
			m.Next()
			return m, current.input.Blur
		}
	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}
func (m Main) View() string {
	current := m.questions[m.index]
	if m.done {
		var output string
		for _, q := range m.questions {
			output += fmt.Sprintf("%s: %s\n", q.question, q.answer)
		}
		return output

	}

	if m.width == 0 {
		return "loading..."
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			current.question,
			m.styles.InputField.Render(current.input.View())))
}

func (m *Main) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func main() {
	fmt.Println("coolmaths!")

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var r1, r2 int

	var questions [3]Question

	for i := 0; i < 3; i++ {
		r1 = r.Intn(10)
		r2 = r.Intn(10)
		questions[i] = newShortQuestion(fmt.Sprintf("What is %d x %d ? ", r1, r2), r1*r2)
	}

	m := New(questions)

	f, err := tea.LogToFile("debug.log", "debug")

	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
