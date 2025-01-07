package pkg

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string // items on the to-do list
	cursor   int      // which to-do list item our cursor is pointing at
	selected int      // which to-do items are selected
}

func NewTeaModel(list []string) (string, error) {
	initModel := Model{
		// Our to-do list is a grocery list
		choices: list,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: -1,
	}
	p := tea.NewProgram(initModel)
	m, err := p.Run()
	if err != nil {
		return "", err
	}
	finalModel, ok := m.(Model)
	if !ok {
		return "", errors.New("model does not implement Model")
	}
	return finalModel.choices[finalModel.selected], nil
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := "Choose an option:\n\n"

	for i, choice := range m.choices {
		cursor := " " // 默认无光标
		if m.cursor == i {
			cursor = ">" // 显示光标
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}
