package interest
import "fmt"
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
        if (balance < 0.000){
        	return float32(3.213)
        } else if (balance >= 0.000 && balance < 1000.000) {
        	return float32(0.5)
        } else if (balance >= 1000 && balance <5000) {
        	return float32(1.621)
        } 
        	return float32(2.475)
    }

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return ((float64(InterestRate(balance))/100)*balance)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.

func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
    // return balance + float64(InterestRate(balance))

}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
// Start with:
// years = 0
// currentBalance = startBalance
// While currentBalance < targetBalance:
// currentBalance = currentBalance × (1 + rate)
// years++
// Return years
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yrs := 0
    curr := balance
    // rate := InterestRate(balance)
    // for curr < targetBalance {
    //     curr = curr * (1+float64(rate/100))
        fmt.Println("case:", curr)
    //     yrs++
    // }
    for (curr < targetBalance) {
        curr += Interest(curr)
        yrs++
    }
    return yrs
}
