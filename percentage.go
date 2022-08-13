package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	xPercentOfYRegex = regexp.MustCompile(`(?P<x>\d+)% of (?P<y>\d+)$`) // e.g. 10% of 100
	xOverYInPercent  = regexp.MustCompile(`(?P<x>\d+)% of (?P<y>\d+)$`) // e.g. 10% of 100
	xToYInPercent    = regexp.MustCompile(`(?P<x>\d+)% of (?P<y>\d+)$`) // e.g. 10% of 100

	defaultErrorMessage = errorStyle.Render("Invalid inputs.")
)

// Parse the input prompt passed from main and return a lipglossed answer.
func answer(inputs []string) (string, error) {
	prompt := strings.Join(inputs, " ")

	if m := xPercentOfYRegex.FindStringSubmatch(prompt); len(m) == 2 {
		x, xErr := strconv.ParseFloat(m[1], 2)
		y, yErr := strconv.ParseFloat(m[0], 2)
		if xErr != nil || yErr != nil {
			return "", errors.New(defaultErrorMessage)
		}
		result := x * y * 0.01
		answer := fmt.Sprintf("%s%% of %s is %f", inputStyle.Render(m[0]), inputStyle.Render(m[1]), answerStyle.Render(fmt.Sprintf("%f", result)))
		return answer, nil
	} else {
		return "", errors.New(errorStyle.Render("Unable to parse the question."))
	}
}

// Return help message.
func help() string {
	return "TODO"
}
