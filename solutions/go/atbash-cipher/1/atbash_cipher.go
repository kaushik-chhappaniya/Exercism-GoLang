package atbashcipher
import (
    "fmt"
    "unicode"
    "strings"
)
func Atbash(s string) string {
    var sb strings.Builder
    str := strings.ToLower(s)
    wordLen := 0
	for _,v := range(str) {
        if unicode.IsLetter(v) {
        	wordLen++
            sb.WriteRune(rune(219-int32(v)))
        } else if unicode.IsNumber(v){
        	wordLen++
            sb.WriteRune(v)
        }
        if wordLen == 5 {
            sb.WriteString(" ")
            wordLen = 0
        }
    }
    result := strings.TrimSpace(sb.String())
    fmt.Println(result)
    return result
}
