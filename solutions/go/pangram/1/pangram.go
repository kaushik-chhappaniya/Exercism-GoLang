package pangram
import (
    "strings"
    "unicode"
)
func IsPangram(input string) bool {
    input = strings.ReplaceAll(strings.ToLower(input), " ","")
	hm := make(map[rune]bool)
    for _,v := range(input){
        if unicode.IsLetter(v) {
        	hm[v] = true
        }
    }
    return len(hm) == 26
}
