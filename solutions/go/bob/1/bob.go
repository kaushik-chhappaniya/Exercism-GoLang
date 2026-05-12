// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob
import (
    "strings"
    "fmt"
    "unicode"
)
 
func Hey(remark string) string {
    question := strings.HasSuffix(strings.TrimSpace(remark), "?")
    silence := strings.TrimSpace(remark) == "" 
 	hasLetter := false
    hasLower := false
    for _, r := range strings.TrimSpace(remark) {
        if unicode.IsLetter(r) {
            hasLetter = true
            if unicode.IsLower(r) {
                hasLower = true
                break
            }
        }
    }  
    yelling := hasLetter && !hasLower 
    
	switch {
        case question && yelling:
        	return fmt.Sprintf("Calm down, I know what I'm doing!")
        case  yelling:
        	 return fmt.Sprintf("Whoa, chill out!")
        case silence:
        	return fmt.Sprintf("Fine. Be that way!")
        case question:
        	return fmt.Sprintf("Sure.")
        default :
        	return fmt.Sprintf("Whatever.")
    }
}
