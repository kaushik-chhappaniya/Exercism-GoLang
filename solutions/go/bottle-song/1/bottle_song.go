package bottlesong
import (
    "fmt"
    "strings"
    )
func Recite(startBottles, takeDown int) []string {
    firstLine := "%s green %s hanging on the wall,"
	secondLine := "And if one green %s should accidentally fall,"
    thirdLine := "There'll be %s green %s hanging on the wall."
    noBottles := "There'll be no green bottles hanging on the wall."
    var hm = map[int]string{
        1:"One",
        2:"Two",
        3:"Three",
        4:"Four",
        5:"Five",
        6:"Six",
        7:"Seven",
        8:"Eight",
        9:"Nine",
        10:"Ten",
        11:"Eleven",
        12:"Twelve",
        13:"Thirteen",
    }
    
    if startBottles < takeDown {
        return []string{""}
    }
    var str []string
    for i:=0; i<takeDown;i++ {
        if i>=1{
            str = append(str, "")
        }
        if val,ok:=hm[startBottles-i]; ok {
            if startBottles-i == 1{
        		str = append(str, fmt.Sprintf(firstLine, val, "bottle"))
        		str = append(str, fmt.Sprintf(firstLine, val, "bottle"))
            } else {
        		str = append(str, fmt.Sprintf(firstLine, val, "bottles"))
        		str = append(str, fmt.Sprintf(firstLine, val, "bottles"))
            }
        }   
                
        if _,ok:=hm[i+1]; ok {
        		str = append(str, fmt.Sprintf(secondLine, "bottle"))
        }
    	bottleFell := startBottles - (i+1)
        if bottleFell == 0 {
        	str = append(str, noBottles)
        }
        if val,ok:=hm[bottleFell]; ok {
            if bottleFell == 1 {
        		str = append(str, fmt.Sprintf(thirdLine, strings.ToLower(val), "bottle"))
            } else {
        		str = append(str, fmt.Sprintf(thirdLine, strings.ToLower(val), "bottles"))
            }
        }
    }	
    return str      
}
