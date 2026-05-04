// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym
import (
    "fmt"
    "strings"
    "unicode"
)
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// l := strings.ReplaceAll(s, "-", " ")
    // k := strings.Fields(l)
	str := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})
	fmt.Println("str:",str)
	var result strings.Builder
    for _, word := range(str){
        if len(word) == 0 {
            continue
        }
        for _, w := range(word ){
            if unicode.IsLetter(w) {
        		result.WriteRune(unicode.ToUpper(w))
        		break
            }
        }
    }
    if len(result.String()) == 0 {
		return ""
    }
    return result.String()
}
