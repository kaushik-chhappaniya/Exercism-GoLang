package luhn
import (
    "strings"
    "unicode"
    "fmt"
)
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) < 2 {
        return false
    }
    fmt.Println("id:", id)

	sum := 0
    double := false
    for i:=len(id)-1; i>=0;i--{
        ch := rune(id[i])
        fmt.Println(i,"--",ch)
        if !unicode.IsDigit(ch){
            return false
        }
        
        digit := int(ch - '0')
        fmt.Println(i,"-->",digit)
        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        		sum += digit
        fmt.Println("sum:",sum)
        	    double = !double
    }
    return sum%10==0
}
