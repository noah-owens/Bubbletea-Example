/*-------------------------------------------------------
|														|
|	Tic-Tac-Toe game created by Noah Owens, utilizes	|
|	skills learned from the BubbleTea basics tutorial	|
|														|
-------------------------------------------------------*/

package main

import (
	"fmt"

	"os"

	tea "github.com/charmbracelet/bubbletea"
)

/*
| model stores application state -- typically a `struct`
*/
type model struct {
	boxes   []string         // placeholder to make UI creation easier
	cursor  int              // which box the cursor is pointing at
	selectX map[int]struct{} // which boxes contain an X
	selectO map[int]struct{} // which boxes contain an O
}

/*
| define inital application state -- can be declared as a variable as well
| but in this case it's defined as a function returning the inital state
*/
func initialModel() model {
	return model{
		// set board length of 9
		// create maps for X's and O's which refer to the box index
		boxes:   []string{"", "", "", "", "", "", "", "", ""},
		selectX: make(map[int]struct{}),
		selectO: make(map[int]struct{}),
	}
}

/*
| Init() method can be utilized to perform inital I/O
| in this case it returns nil, as no I/O is required.
*/
func (m model) Init() tea.Cmd {
	return nil
}

/*
| Update() method listens for Msg and changes the application
| state accordingly. `KeyMsg` messages are sent to Update()
| automatically
*/
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// nested switch statement listens for msg, verifies type as keystrokes
	// then determines which key and updates accordingly.
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// exit program
		case "ctrl+c", "q":
			return m, tea.Quit

		// move cursor 'up'
		case "up", "left", "a", "w":
			if m.cursor > 0 {
				m.cursor--
			}

		// move cursor 'down'
		case "down", "right", "d", "s":
			if m.cursor < len(m.boxes) {
				m.cursor++
			}

		// alternate selection between X's and O's
		case "enter", " ":
			if len(m.selectX) == 0 || len(m.selectX) <= len(m.selectO) {
				m.selectX[m.cursor] = struct{}{}
			} else {
				m.selectO[m.cursor] = struct{}{}
			}
		}
	}

	// return updated model with no commands
	return m, nil
}

/*
| View() method returns a string to be used as the UI
*/
func (m model) View() string {
	// header
	s := "Tic-Tac-Toe\n\n"

	// iterate across boxes
	for i, box := range m.boxes {

		// cursor pointing at this box?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		// is choice marked with X or O?
		checked := " "
		if _, ok := m.selectX[i]; ok {
			checked = "x"
		}
		if _, ok := m.selectO[i]; ok {
			checked = "o"
		}

		// render the grid
		if i > 0 && i%3 == 2 {
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, box)
		} else {
			s += fmt.Sprintf("%s [%s] %s", cursor, checked, box)
		}
	}

	// footer
	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Sorry, there's been an error: %v", err)
		os.Exit(1)
	}
}
