package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.RoundedBorder()).Padding(1).Width(80)
	return s
}

type Main struct {
	styles    *Styles
	title1    string
	title2    string
	index     int
	questions [30]Question
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

func InitializeMainScreen(questions [30]Question) *Main {
	styles := DefaultStyles()
	title1 := "C O O L M A T H S"
	title2 := "Learn, Play and Enjoy Maths"
	return &Main{styles: styles, title1: title1, title2: title2, questions: questions}
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
		case "r":
			m.questions = createQuestions()
			m.index = 0
			m.done = false
			current = &m.questions[m.index]
			return m, current.input.Blur
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			}
			current.answer = current.input.Value()
			m.Next()
			return m, current.input.Blur
		}
	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}
func (m Main) View() string {
	var rightAnswerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#b1ff9c"))

	var wrongAnswerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#fa4d4d"))

	current := m.questions[m.index]
	if m.done {
		columns := []table.Column{
			{Title: "Question", Width: 10},
			{Title: "Problem", Width: 30},
			{Title: "Your Answer", Width: 15},
			{Title: "Correct Answer", Width: 15},
			{Title: "Feedback", Width: 20},
		}
		var rows []table.Row
		count := 0
		var ans int
		var response string

		for qn, q := range m.questions {
			if len(strings.TrimSpace(q.answer)) == 0 {
				response = "Doh!"
			} else {

				ans, _ = strconv.Atoi(q.answer)
				if ans == q.expected {
					response = "You got it!"
					count += 1
				} else {
					response = "Doh!"
				}

			}

			rows = append(rows, table.Row{
				fmt.Sprintf("%d", qn+1),
				fmt.Sprintf("%s", q.question),
				fmt.Sprintf("%s", q.answer),
				fmt.Sprintf("%d", q.expected),
				response,
			})
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(false),
			table.WithHeight(30),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)

		t.SetStyles(s)
		var final string

		percent := (float32(count) / 30.0) * 100.0

		if percent == 100.0 {
			final += "\n\n" + rightAnswerStyle.Render("You Won!")
		} else {
			final += "\n\n" + wrongAnswerStyle.Render("Oh Oh. You will more need practice")
		}

		final += fmt.Sprintf("\n\nYou got %2d out 30 correct. You scored %0.2f %%", count, float32(percent))
		final += fmt.Sprintf("\n\nPress q to exit or Press r to play again.")

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Top,
			tableStyle.Render(t.View())+final)

	}

	if m.width == 0 {
		return "loading..."
	}

	var style1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Background(lipgloss.Color("#8cf5d2")).
		Padding(2).
		MarginBottom(2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		BorderBackground(lipgloss.Color("#8cf5d2"))

	var style2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("428")).
		Padding(2).
		BorderForeground(lipgloss.Color("428"))

	title := lipgloss.JoinVertical(lipgloss.Center, style1.Render(m.title1), style2.Render(m.title2))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			lipgloss.JoinVertical(
				lipgloss.Left,
				current.question,
				m.styles.InputField.Render(current.input.View()))))
}

func (m *Main) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func createQuestions() [30]Question {
	var r1, r2 int

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var questions [30]Question

	for i := 0; i < 30; i++ {
		r1 = r.Intn(10)
		r2 = r.Intn(10)
		questions[i] = newShortQuestion(fmt.Sprintf("What is %d x %d ? ", r1, r2), r1*r2)
	}
	return questions
}

func main() {
	questions := createQuestions()
	m := InitializeMainScreen(questions)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
