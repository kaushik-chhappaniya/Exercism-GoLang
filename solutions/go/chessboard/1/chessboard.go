package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var sum int = 0
    for _,v := range cb[file]{
        if v {
            sum ++
        }
    }
    return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    var sum int = 0
    if rank > 8 || rank <= 0 {
        return 0
    }
    for _, val := range cb {
        if val[rank-1] {
          sum ++
        }
    }
    return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    // length of ranks * number of files in each rank 
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var sum int = 0
    for val := range cb {
        sum += CountInFile(cb, val)
    }
    return sum
}
