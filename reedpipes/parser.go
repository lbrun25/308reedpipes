package reedpipes

import (
	"errors"
	"os"
	"strconv"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"

	invalidValues = "Invalid values.\n"
	negativeValues = "One or more values are negative.\n"
	nLessThan2 = "n must be greater than 2"

	maxArg = 6
	minArg = 6
)

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if arg == "-h" {
			return true
		}
	}
	return false
}

// CheckArgs - check the user's args
func CheckArgs() error {
	argsWithoutProg := os.Args[1:]

	// Check the number of arguments
	if len(argsWithoutProg) < minArg {
		return errors.New(notEnoughArgs)
	}
	if len(argsWithoutProg) > maxArg {
		return errors.New(tooManyArgs)
	}

	// Retrieve values
	var values []float64

	for _, arg := range argsWithoutProg {
		value, err := strconv.ParseFloat(arg, 64); if err != nil {
			return errors.New(invalidValues)
		}
		if value < 0 {
			return errors.New(negativeValues)
		}
		values = append(values, value)
	}
	r0 = values[0]
	r5 = values[1]
	r10 = values[2]
	r15 = values[3]
	r20 = values[4]
	n = values[5]

	if n < 2 {
		return errors.New(nLessThan2)
	}

	return nil
}