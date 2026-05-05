package armstrongnumbers
import (
    "math"
    "strconv"
)
func IsNumber(n int) bool {
    var cp int
    cp = n
    res := 0
    length := len(strconv.Itoa(n))
	for n > 0{
        num := n % 10
        n /= 10
        res += int(math.Pow(float64(num) , float64(length)))
    }
    return res == cp
}
