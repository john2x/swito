package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
)

const (
	inputColor  = lipgloss.Color("5") // magenta
	answerColor = lipgloss.Color("6") // cyan
	errorColor  = lipgloss.Color("9") // red
)

var (
	inputStyle  = lipgloss.NewStyle().Bold(true).Foreground(inputColor)
	answerStyle = lipgloss.NewStyle().Bold(true).Foreground(answerColor)
	errorStyle  = lipgloss.NewStyle().Foreground(errorColor)
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(errorStyle.Render("What do you want to know?"))
		os.Exit(1)
	}

	answer, err := answer(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(answer)
}
