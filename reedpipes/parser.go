package reedpipes

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrTooManyArgs = errors.New("there are too many arguments")
	ErrNotEnoughArgs = errors.New("there are not enough arguments")
	ErrInvalidValues = errors.New("invalid values")
	ErrNegativeValues = errors.New("one or more values are negative")
	ErrNLessThan2 = errors.New("n must be greater than 2")
)

const (
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
		return ErrNotEnoughArgs
	}
	if len(argsWithoutProg) > maxArg {
		return ErrTooManyArgs
	}

	// Retrieve values
	var values []float64

	for _, arg := range argsWithoutProg {
		value, err := strconv.ParseFloat(arg, 64); if err != nil {
			return ErrInvalidValues
		}
		if value <= 0 {
			return ErrNegativeValues
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
		return ErrNLessThan2
	}

	return nil
}