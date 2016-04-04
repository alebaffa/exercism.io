package queenattack

import "errors"

func CanQueenAttack(whiteQueen, blackQueen string) (bool, error) {

	if !isInBoard(whiteQueen) || !isInBoard(blackQueen) || whiteQueen == blackQueen {
		return false, errors.New("off the board")
	}

	distanceInRow, distanceInCol := calculateDistances(whiteQueen, blackQueen)
	return (distanceInRow == 0 || distanceInCol == 0 || distanceInRow == distanceInCol), nil
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
	return int(whiteQueen[0]) - int(blackQueen[0]), int(whiteQueen[1]) - int(blackQueen[1])
}
