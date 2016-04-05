package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a given position of the chessboard
func Square(position int) (uint64, error) {
	if position < 1 || position > 64 {
		return 0, errors.New("position not valid")
	}
	grains := uint64(math.Pow(2, float64(position-1)))
	return grains, nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return uint64(math.Pow(2, float64(64))) - 1
}
