package rnatranscription
import "strings"

func ToRNA(dna string) string {
    if dna == "" {
        return ""
    }
    replacements := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	var result strings.Builder
    for _, v:=range(dna){
        if replaceChar, exists := replacements[v]; exists {
            result.WriteRune(replaceChar)
        } else {
            result.WriteRune(v)
        }
    }
    return result.String()
}
