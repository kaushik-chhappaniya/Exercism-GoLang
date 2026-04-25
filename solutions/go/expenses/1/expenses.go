package expenses
import "fmt"
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	if predicate == nil {
        panic("NA")
    }
    var arr []Record
    for _,v := range in {
        if predicate(v){
        	arr = append(arr, v)
        }
    }
    return arr
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	from := p.From
    to := p.To
    return func(r Record) bool{
        return (r.Day >= from && r.Day <= to) 
    } 
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        return c == r.Category
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    var sum float64
    fn := ByDaysPeriod(p)
	for _, v := range in {
        if fn(v){
            sum += v.Amount
        }
    }
    return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64
    var check int
    daysFunc := ByDaysPeriod(p)
    categoryFunc := ByCategory(c)
    for _,v:=range in{
        if categoryFunc(v) {
        	if daysFunc(v){
            	sum += v.Amount
            }     
        } else {
         check++   
         }    	
    }
    if check == len(in) {
       return 0, fmt.Errorf("unknown category entertainment")
    }
    return sum, nil
}
