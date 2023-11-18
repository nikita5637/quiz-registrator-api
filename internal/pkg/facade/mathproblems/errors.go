package mathproblems

import "errors"

const (
	// ReasonMathProblemAlreadyExists ...
	ReasonMathProblemAlreadyExists = "MATH_PROBLEM_ALREADY_EXISTS"
	// ReasonMathProblemNotFound ...
	ReasonMathProblemNotFound = "MATH_PROBLEM_NOT_FOUND"
)

var (
	// ErrMathProblemAlreadyExists ...
	ErrMathProblemAlreadyExists = errors.New("math problem already exists")
	// ErrMathProblemNotFound ...
	ErrMathProblemNotFound = errors.New("math problem not found")
)
