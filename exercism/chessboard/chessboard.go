package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileInBoard, exists := cb[file]
	if !exists {
		return 0
	}

	acc := 0

	for _, value := range fileInBoard {
		if value {
			acc += 1
		}
	}

	return acc
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	acc := 0
	adjRank := rank - 1

	if adjRank < 0 || adjRank > 7 {
		return acc
	}

	for _, file := range cb {
		if file[adjRank] {
			acc += 1
		}
	}

	return acc
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// return 64 would do but let's do it properly
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	acc := 0
	for _, file := range cb {
		for _, isOccupied := range file {
			if isOccupied {
				acc++
			}
		}
	}

	return acc
}
