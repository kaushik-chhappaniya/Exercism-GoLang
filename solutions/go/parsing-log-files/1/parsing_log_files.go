package parsinglogfiles

import (
	"fmt"
	"regexp"
    "strings"
)
func IsValidLine(text string) bool {
    re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)    
}

func SplitLogLine(text string) []string {
    //Split using only given symbols
	re := regexp.MustCompile(`<[=_~*-]*>`)
    split := re.Split(text, -1)
    var s []string
	for _, v := range split {
		s = append(s, v)
	}
    return s
}

func CountQuotedPasswords(lines []string) int {
	re:= regexp.MustCompile("[^`]password[^`]")
    var total int =0
    for _,line := range(lines){
    	arr := re.FindAllString(strings.ToLower(line),-1)
        total += len(arr)
    }
    return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S*)`)
    var arr []string 
    for _, line := range(lines) {
        find := re.FindStringSubmatch(line)
        s:=line
        if len(find) > 0 {
        	s = fmt.Sprintf("[USR] %s %s", find[1], line)
        }
        	arr = append(arr, s)
    }
	return arr
}
