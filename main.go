package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []emoji
	selected int
	search   string
}

func initialModel() model {
	var head []emoji
	i := 0
	for _, pair := range Emojies {
		if i == 5 {
			break
		}
		head = append(head, pair)
		i++
	}

	return model{
		choices:  head,
		selected: 0,
		search:   "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "escape":
			return m, tea.Quit

		case "down":
			if m.selected-1 < len(m.choices) {
				m.selected++
			}

		case "up":
			if m.selected > 0 {
				m.selected--
			}

		case "enter":
			emoji := m.choices[m.selected]
			err := clipboard.WriteAll(emoji.emoji)
			if err == nil {
				fmt.Printf("✅ Successfuly copied %s \"%s\" to clipboard!\n", emoji.emoji, emoji.title)
			}
			return m, tea.Quit

		case "backspace":
			if len(m.search) > 0 {
				m.search = m.search[:len(m.search)-1]
			}

		default:
			m.search = m.search + string(msg.Runes)
		}

	}

	var head []emoji
	found := 0
	for _, pair := range Emojies {
		if found == 5 {
			break
		}
		if strings.Contains(strings.ToLower(pair.title), strings.ToLower(m.search)) {
			head = append(head, pair)
			found++
		}
	}
	m.choices = head

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Search: %s\n", m.search)

	for i, choice := range m.choices {
		cursor := " "
		if m.selected == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, choice.emoji, choice.title)
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Errorf("Failed to run the program: %v\n", err)
		os.Exit(1)
	}
}
