/*
Package grains implements logic for calculating
number of grains of wheat on a chessboard.
*/
package grains

import (
	"github.com/pkg/errors"
	"math"
)

// Square returns doubled number of grains of wheat based
// on the given number of chessboard square which can be from
// 1 to 64, otherwise will return an error
func Square(i int) (uint64, error) {
	if i >= 1 && i <= 64 {
		return uint64(math.Pow(2, float64(i-1))), nil
	}
	return 0, errors.Errorf("i should be between 1 and 64 inclusive, was %d", i)
}

// Total returns the total amount of grains of wheat on the chessboard
func Total() uint64 {
	if sq, err := Square(64); err == nil {
		return sq
	}
	return 0
}
