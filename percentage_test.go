package main

import "testing"

func TestAnswer(t *testing.T) {
	args := []string{"10%", "of", "90"}
	answer, err := answer(args)
	if err != nil {
		t.Errorf(err.Error())
	}
	if answer != "10% of 90 is 9" {
		t.Errorf(answer)
		t.Errorf("Answer is incorrect")
	}
}
