package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nCows int) (float64, error) {

	totalFodder, err := fc.FodderAmount(nCows)

	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return totalFodder * factor / float64(nCows), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nCows int) error {
	if nCows < 0 {
		return &InvalidCowsError{
			nCows:   nCows,
			message: "there are no negative cows",
		}
	} else if nCows == 0 {
		return &InvalidCowsError{
			nCows:   nCows,
			message: "no cows don't need food",
		}
	}
	return nil
}

type InvalidCowsError struct {
	nCows   int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.nCows, e.message)
}
