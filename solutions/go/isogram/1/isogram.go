package isogram
import "strings"
func IsIsogram(word string) bool {
	hm:=make(map[rune]bool)
    word = strings.ToLower(word)
    for _, v := range(word){
        if v != '-' && v != ' '{
            if  hm[v] {
                return false
            } 
            hm[v]=true
        }
    }
        return true
}
