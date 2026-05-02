package differenceofsquares

func SquareOfSum(n int) int {
    if n == 1 {
        return 1
    }
    sum := (n*(1+n))/2
    return sum*sum
}

func SumOfSquares(n int) int {
    if n == 1 {
        return 1
    }
    sumOfSq := (n*(n+1)*((2*n)+1))/6
	return sumOfSq
    
}

func Difference(n int) int {
	return SquareOfSum(n)- SumOfSquares(n) 
}
