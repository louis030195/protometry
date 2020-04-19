package protometry

import (
	"errors"
)

var (
	// ErrDivisionByZero ...
	ErrDivisionByZero = errors.New("Division by zero")
	// ErrVector3otSameSize ...
	ErrVector3otSameSize = errors.New("Vectors are not the same size")
	// ErrVectorInvalidDimension ...
	ErrVectorInvalidDimension = errors.New("Vectors' dimensions are not of the expected size")
	// ErrVectorInvalidIndex ...
	ErrVectorInvalidIndex = errors.New("Invalid index")
)
