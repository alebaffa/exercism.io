package queenattack

import "errors"

func CanQueenAttack(whiteQueen, blackQueen string) (bool, error) {

	if !isInBoard(whiteQueen) || !isInBoard(blackQueen) {
		return false, errors.New("off the board")
	}
	if whiteQueen == blackQueen {
		return false, errors.New("queens in the same position")
	}

	distanceInRow, distanceInCol := calculateDistances(whiteQueen, blackQueen)
	return (distanceInRow == 0 || distanceInCol == 0 || distanceInRow == distanceInCol || distanceInRow == -distanceInCol), nil
}

func isInBoard(queen string) bool {
	if len(queen) != 2 {
		return false
	}
	rowInBoard := ('a' <= queen[0] && queen[0] <= 'h')
	colInBoard := ('1' <= queen[1] && queen[1] <= '8')
	return rowInBoard && colInBoard
}

func calculateDistances(whiteQueen, blackQueen string) (int, int) {
	rowDistance := int(whiteQueen[0]) - int(blackQueen[0])
	colDistance := int(whiteQueen[1]) - int(blackQueen[1])
	return rowDistance, colDistance
}
