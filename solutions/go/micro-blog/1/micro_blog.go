package microblog

func Truncate(phrase string) string {
	arr := []rune(phrase)
    if len(arr) <=5 {
        return phrase
    }
    return string(arr[:5])
}
