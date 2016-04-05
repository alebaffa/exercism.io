package queenattack

import "errors"

// CanQueenAttack calculates the positions of the two queens in input and defines if the first queen can attack the second one.
func CanQueenAttack(whiteQueen, blackQueen string) (bool, error) {

	if !isValid(whiteQueen) || !isValid(blackQueen) {
		return false, errors.New("coordinates are not valid")
	}

	if !isInBoard(whiteQueen) || !isInBoard(blackQueen) {
		return false, errors.New("off the board")
	}

	if whiteQueen == blackQueen {
		return false, errors.New("queens in the same position")
	}

	rowDist, colDist := distances(whiteQueen, blackQueen)
	return (onSameRowOrColumn(rowDist, colDist) || onSameDiagonal(rowDist, colDist)), nil
}

func isInBoard(queen string) bool {
	rowInBoard := ('a' <= queen[0] && queen[0] <= 'h')
	colInBoard := ('1' <= queen[1] && queen[1] <= '8')
	return rowInBoard && colInBoard
}

func isValid(queen string) bool {
	if len(queen) != 2 {
		return false
	}
	return true
}

func distances(whiteQueen, blackQueen string) (int, int) {
	rowDistance := int(whiteQueen[0]) - int(blackQueen[0])
	colDistance := int(whiteQueen[1]) - int(blackQueen[1])
	return rowDistance, colDistance
}

func onSameRowOrColumn(rowDist, colDist int) bool {
	return rowDist == 0 || colDist == 0
}

func onSameDiagonal(rowDist, colDist int) bool {
	return rowDist == colDist || rowDist == -colDist
}
