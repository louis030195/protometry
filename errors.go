package protometry

import (
	"errors"
)

var (
	// ErrDivisionByZero ...
	ErrDivisionByZero = errors.New("Division by zero")
	// ErrVectorNotSameSize ...
	ErrVectorNotSameSize = errors.New("Vectors are not the same size")
	// ErrVectorInvalidDimension ...
	ErrVectorInvalidDimension = errors.New("Vectors' dimensions are not of the expected size")
	// ErrVectorInvalidIndex ...
	ErrVectorInvalidIndex = errors.New("Invalid index")
)
