package hamming
import "fmt"
func Distance(a, b string) (int, error) {
	l1 := len(a)
    l2 := len(b)
switch {
    case (l1 == 0 && l2 == 0) :
        return 0, nil
    case (l1 <= 0 || l2 <= 0) :
		return 0, fmt.Errorf("error")
    case l1 != l2:
        return 0, fmt.Errorf("error")
   default:{
            s:=0 
        	min1 := min(l1, l2)
        	for i := 0; i < min1; i++ {
        		if a[i] != b[i] {
        			s++
        		}
        	}
        	return s, nil
   		}     
	}
}
