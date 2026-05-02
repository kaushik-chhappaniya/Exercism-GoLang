package thefarm
import (
    "errors"
    "fmt"
    )



// TODO: define the 'DivideFood' function
// FodderCalculator as mentioned is implemented somewhere else so we just need to use it.
// No wrapping to error interface by any means like fmt.Errorf or errors.New nothing
func DivideFood(fodderCalculator FodderCalculator, cows int)(float64, error){
    
	fdFc, err := fodderCalculator.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    
	ftFc, err := fodderCalculator.FatteningFactor()
    if err!= nil {
        return 0, err
    }
    food:=(fdFc*ftFc)/float64(cows)
    return  food, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error){
    if cows <= 0 {
    	return 0, errors.New("invalid number of cows")
        }  else {
        val, err := DivideFood(fodderCalculator, cows)
        if err != nil {
            return 0, err
        }
        return val, nil
    }
}

// Defining custom struct that can be used as caller for error method
type InvalidCowsError struct {
    Cows    int
    Message string
}

// the struct that is implementing Error() func is automatically a custom error. Error interface is already implemented in the base package
func (e *InvalidCowsError) Error() string{
    return fmt.Sprintf("%d cows are invalid: %s", e.Cows, e.Message)
}

// TODO: define the 'ValidateNumberOfCows' function
// Here we're using the invalid custom error created by passing the required data to t.
func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{
            Cows:    cows,
            Message: "there are no negative cows",
        }
    }
    if cows == 0 {
        return &InvalidCowsError{
            Cows:    cows,
            Message: "no cows don't need food",
        }
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
