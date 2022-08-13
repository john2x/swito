package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	xPercentOfYRegex = regexp.MustCompile(`^(?P<x>([0-9]*[.])?[0-9]+)% of (?P<y>[-]?([0-9]*[.])?[0-9]+)$`) // e.g. 10% of 100
	xOverYInPercent  = regexp.MustCompile(`(?P<x>\d+)% of (?P<y>\d+)$`)                                    // e.g. 10% of 100
	xToYInPercent    = regexp.MustCompile(`(?P<x>\d+)% of (?P<y>\d+)$`)                                    // e.g. 10% of 100

	defaultErrorMessage = errorStyle.Render("Invalid inputs.")
)

// Parse the input prompt passed from main and return a lipglossed answer.
func answer(inputs []string) (string, error) {
	prompt := strings.Join(inputs, " ")

	if m := xPercentOfYRegex.FindStringSubmatch(prompt); len(m) == 5 {
		x, xErr := strconv.ParseFloat(m[1], 2)
		y, yErr := strconv.ParseFloat(m[3], 2)
		if xErr != nil || yErr != nil {
			return "", errors.New(defaultErrorMessage)
		}
		result := x * y * 0.01
		answer := fmt.Sprintf("%s%% of %s is %s", inputStyle.Render(m[1]), inputStyle.Render(m[3]), answerStyle.Render(fmt.Sprintf("%.2f", result)))
		return answer, nil
	} else {
		return "", errors.New(errorStyle.Render("Unable to parse the question."))
	}
}

// Return help message.
func help() string {
	return "TODO"
}
