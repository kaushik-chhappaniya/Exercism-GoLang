package scrabblescore
import (
    "strings"
    "fmt"
    )
func Score(word string) int {
    sum :=0
    for _, v:=range(strings.ToUpper(word)) {
        sum += getVal(v)
    }
    return sum
}

func getVal(c rune) int {
    switch {
        // case 'A', 'E', 'I', 'O', 'U',  'L', 'N', 'R', 'S', 'T':
        // return 1
        case strings.ContainsRune("AEIOULNRST", c):
        	return 1
        case strings.ContainsRune("DG", c):
        	return 2
        case strings.ContainsRune("BCMP", c):
        	return 3
        case strings.ContainsRune("FHVWY", c):
            return 4
        case strings.ContainsRune("K", c):
        	return 5
        case strings.ContainsRune("JX", c):
        	return 8
        case strings.ContainsRune("QZ", c):
        	return 10
        default:
        	fmt.Println(string(c))
        	return 0
    }
}