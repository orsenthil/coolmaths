package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Blink() tea.Msg
	Blur() tea.Msg
	Focus() tea.Cmd
	SetValue(string)
	Value() string
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

type ShortAnswerField struct {
	textinput textinput.Model
}

func NewShortAnswerField() *ShortAnswerField {
	a := ShortAnswerField{}

	model := textinput.New()
	model.Placeholder = "Your answer here"
	model.Focus()

	a.textinput = model
	return &a
}

func (a *ShortAnswerField) Blink() tea.Msg {
	return textinput.Blink()
}
func (a *ShortAnswerField) Init() tea.Msg {
	return nil
}

func (a *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	a.textinput, cmd = a.textinput.Update(msg)
	return a, cmd
}

func (a *ShortAnswerField) View() string {
	return a.textinput.View()
}

func (a *ShortAnswerField) Focus() tea.Cmd {
	return a.textinput.Focus()
}

func (a *ShortAnswerField) SetValue(s string) {
	a.textinput.SetValue(s)
}
func (a *ShortAnswerField) Blur() tea.Msg {
	return a.textinput.Blur
}

func (a *ShortAnswerField) Value() string {
	return a.textinput.Value()
}
