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
	windowStart int
	width int
}

func initialModel() model {
    return model{
        choices:     Emojies,
        selected:    0,
        search:      "",
        windowStart: 0,
    }
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

	case tea.WindowSizeMsg:
        m.width = msg.Width

    case tea.KeyMsg:
        switch msg.String() {

        case "escape", "ctrl+c":
            return m, tea.Quit

        case "down":
            if m.selected < len(m.choices)-1 {
                m.selected++
            }

        case "up":
            if m.selected > 0 {
                m.selected--
            }

        case "enter":
            if len(m.choices) == 0 {
                return m, nil
            }
            emoji := m.choices[m.selected]
            err := clipboard.WriteAll(emoji.emoji)
            if err == nil {
                fmt.Printf("✅ Successfully copied %s \"%s\" to clipboard!\n", emoji.emoji, emoji.title)
            }
            return m, tea.Quit

        case "backspace":
            if len(m.search) > 0 {
                m.search = m.search[:len(m.search)-1]
            }

		case " ":
			m.search += " "

        default:
            if msg.Type == tea.KeyRunes {
                m.search = m.search + string(msg.Runes)
            }
        }
    }

	var allMatches []emoji
    for _, pair := range Emojies {
        if strings.Contains(strings.ToLower(pair.title), strings.ToLower(m.search)) {
            allMatches = append(allMatches, pair)
        }
    }
    m.choices = allMatches

	if m.selected >= len(m.choices) {
        m.selected = len(m.choices) - 1
        if m.selected < 0 {
            m.selected = 0
        }
    }

	if m.selected < m.windowStart {
		m.windowStart = m.selected
    } else if m.selected >= m.windowStart+5 {
		m.windowStart = m.selected - 4
    }

	if len(m.choices) == 0 {
        m.windowStart = 0
    }

    return m, nil
}

func (m model) View() string {
    s := fmt.Sprintf("Search: %s\n", m.search)

	end := m.windowStart + 5
    if end > len(m.choices) {
        end = len(m.choices)
    }

	termWidth := m.width
    if termWidth == 0 {
        termWidth = 80 
    }

	for i := m.windowStart; i < end; i++ {
        choice := m.choices[i]
        cursor := "  "
        if m.selected == i {
            cursor = "> "
        }
        title := choice.title
        
        // Calculate max allowed length for the title.
        // We subtract 10 to reserve space for the cursor, emoji, spaces, and margin.
        maxTitleLen := termWidth - 10 
        
        // Safely truncate string using runes
        if maxTitleLen > 0 {
            titleRunes := []rune(title)
            if len(titleRunes) > maxTitleLen {
                title = string(titleRunes[:maxTitleLen-3]) + "..."
            }
        }

        s += fmt.Sprintf("%s%s %s\n", cursor, choice.emoji, title)
    }

    s += fmt.Sprintf("\n(Showing %d of %d results)\n", end-m.windowStart, len(m.choices))
    return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Errorf("Failed to run the program: %v\n", err)
		os.Exit(1)
	}
}
