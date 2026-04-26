package collatzconjecture
import "fmt"
func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
    return 0, fmt.Errorf("input must be a positive integer")
	}
    var steps, num int
    num = n
    for {
        if num == 1{
        	return steps, nil
        }
        if num % 2 == 0 {
            num = num/2
        } else {
    num = num*3 + 1
}
 
        steps++
    }
    return steps,nil
}
